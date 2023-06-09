package azure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"

	credreqv1 "github.com/openshift/cloud-credential-operator/pkg/apis/cloudcredential/v1"
	azureclients "github.com/openshift/cloud-credential-operator/pkg/azure"
	"github.com/openshift/cloud-credential-operator/pkg/cmd/provisioning"

	uuid "github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// CreateManagedIdentitiesOpts are azureOptions necessary for creating user-assigned managed identities
	CreateManagedIdentitiesOpts = azureOptions{}

	secretManifestTemplate = `apiVersion: v1
stringData:
  azure_client_id: %s
  azure_tenant_id: %s
  azure_region: %s
  azure_subscription_id: %s
  azure_federated_token_file: %s
kind: Secret
metadata:
  name: %s
  namespace: %s
type: Opaque`

	ingressCredentialRequestName = "openshift-ingress-azure"
)

// createManagedIdentity creates a user-assigned managed identity for the provided CredentialsRequest
// with name "<name>-<CredentialsRequest.Spec.SecretRef.Namespace>-<CredentialsRequest.Spec.SecretRef.Name>",
// eg "mycluster-openshift-machine-api-azure-cloud-credentials".
//
// The user-assigned managed identity will be assigned pre-existing Azure roles as specified within
// CredentialsRequest.Spec.ProviderSpec.RoleBindings. Role assignment will be scoped within the resource
// groups provided as scopingResourceGroupNames.
//
// Roles which are assigned to pre-existing user-assigned managed identities found to exist for
// processed CredentialsRequests will be removed if not enumerated within the CredentialsRequest
// to idempotently align with the desired state.
//
// A federated identity credential will be created for each service account enumerated within
// CredentialsRequest.Spec.ServiceAccountNames.
//
// A secret containing user-assigned managed identity details will be written to the outputDir
// once the user-assigned managed identity is created and configured.
func createManagedIdentity(client *azureclients.AzureClientWrapper, name, resourceGroupName, subscriptionID, region, issuerURL, outputDir string, scopingResourceGroupNames []string, resourceTags map[string]string, credentialsRequest *credreqv1.CredentialsRequest, dryRun bool) error {
	// Write dummy secrets with blank clientID and tenantID when doing a dry run.
	if dryRun {
		writeCredReqSecret(credentialsRequest, outputDir, "", "", subscriptionID, region)
		return nil
	}

	// Create user-assigned managed identity with name "name-targetNamespace-targetSecretName"
	// Azure resources can't have a name longer than 128 characters
	managedIdentityName := fmt.Sprintf("%s-%s-%s", name, credentialsRequest.Spec.SecretRef.Namespace, credentialsRequest.Spec.SecretRef.Name)
	shortenedManagedIdentityName := provisioning.ShortenName(managedIdentityName, 128)
	userAssignedManagedIdentity, err := ensureUserAssignedManagedIdentity(client, shortenedManagedIdentityName, resourceGroupName, region, resourceTags)
	if err != nil {
		return err
	}

	// Decode CredentialsRequest.Spec.ProviderSpec.RoleBindings from Azure CredentialsRequest
	crProviderSpec := &credreqv1.AzureProviderSpec{}
	if credentialsRequest.Spec.ProviderSpec != nil {
		codec, err := credreqv1.NewCodec()
		if err != nil {
			return err
		}
		err = codec.DecodeProviderSpec(credentialsRequest.Spec.ProviderSpec, crProviderSpec)
		if err != nil {
			return fmt.Errorf("error decoding provider spec from CredentialsRequest: %w", err)
		}
	}

	// List role assignments by the user-assigned managed identity principal ID
	// This list of role assignments are roles which are assigned to the user-assigned managed identity
	existingRoleAssignments := []*armauthorization.RoleAssignment{}
	listRoleAssignments := client.RoleAssignmentClient.NewListPager(
		&armauthorization.RoleAssignmentsClientListOptions{
			Filter: to.Ptr(fmt.Sprintf("principalId eq '%s'", *userAssignedManagedIdentity.Properties.PrincipalID)),
		},
	)
	for listRoleAssignments.More() {
		pageResponse, err := listRoleAssignments.NextPage(context.Background())
		if err != nil {
			return err
		}
		existingRoleAssignments = append(existingRoleAssignments, pageResponse.RoleAssignmentListResult.Value...)
	}

	// shouldExistRoleAssignments are role assignments we validated should exist or were created.
	// shouldExistRoleAssignments are used to determine which role assignments need to be deleted as
	// compared to the roleAssignments listed above.
	shouldExistRoleAssignments := []*armauthorization.RoleAssignment{}

	// Assign requested roles to the user-assigned managed identity
	// Role assignment will be scoped to the resource group identified by scopingResourceGroupName
	for _, roleBinding := range crProviderSpec.RoleBindings {
		//log.Printf("processing role %v", roleBinding.Role)
		for _, scopingResourceGroupName := range scopingResourceGroupNames {
			scope := "/subscriptions/" + subscriptionID + "/resourceGroups/" + scopingResourceGroupName
			// Get Azure role definition for the role name (roleBinding.Role)
			roleDefinition, err := getRoleDefinitionByRoleName(client, roleBinding.Role, subscriptionID)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to get role definition for role %s", roleBinding.Role))
			}

			// Determine if the role definition's ID is already assigned to the user-assigned managed identity
			// at the specified scope
			roleAssignmentExists := false
			for _, roleAssignment := range existingRoleAssignments {
				if *roleDefinition.Properties.RoleName == roleBinding.Role && *roleAssignment.Properties.RoleDefinitionID == *roleDefinition.ID && *roleAssignment.Properties.Scope == scope {
					roleAssignmentExists = true
					log.Printf("Found existing role assignment %s for user-assigned managed identity with principal ID %s at scope %s", roleBinding.Role, *userAssignedManagedIdentity.Properties.PrincipalID, scope)
					shouldExistRoleAssignments = append(shouldExistRoleAssignments, roleAssignment)
					break
				}
			}

			if !roleAssignmentExists {
				// Assign role to identity at scope
				roleAssignment, err := createRoleAssignment(
					client,
					*userAssignedManagedIdentity.Properties.PrincipalID,
					*roleDefinition.ID,
					roleBinding.Role,
					scope,
					subscriptionID,
				)
				if err != nil {
					return errors.Wrapf(err, "failed to assign role %s to user-assigned managed identity", roleBinding.Role)
				}
				shouldExistRoleAssignments = append(shouldExistRoleAssignments, roleAssignment)
			}
		}
	}

	//log.Printf("there are %v existing role assignments", len(existingRoleAssignments))
	//log.Printf("there are %v role assignments that should exist", len(shouldExistRoleAssignments))

	for _, existingRoleAssignment := range existingRoleAssignments {
		found := false
		for _, shouldExistRoleAssignment := range shouldExistRoleAssignments {
			if shouldExistRoleAssignment.Name == existingRoleAssignment.Name {
				found = true
			}
		}
		if !found {
			roleDefinition, err := getRoleDefinitionByID(client, *existingRoleAssignment.Properties.RoleDefinitionID)
			if err != nil {
				return errors.Wrapf(err, "failed to get role definition with role definition ID %s", *existingRoleAssignment.Properties.RoleDefinitionID)
			}
			err = deleteRoleAssignment(client,
				*userAssignedManagedIdentity.Properties.PrincipalID,
				*existingRoleAssignment.Name,
				*roleDefinition.Properties.RoleName,
				*existingRoleAssignment.Properties.Scope,
				subscriptionID,
			)
			if err != nil {
				return errors.Wrapf(err, "failed to unassign role with ID %s from user-assigned managed identity", *roleDefinition.Properties.RoleName)
			}
		}
	}

	// Create a federated identity credential for every service account enumerated in the CredentialsRequest
	for _, serviceAccountName := range credentialsRequest.Spec.ServiceAccountNames {
		err := createFederatedIdentityCredential(client, shortenedManagedIdentityName, issuerURL, credentialsRequest.Spec.SecretRef.Namespace, serviceAccountName, resourceGroupName)
		if err != nil {
			return err
		}
	}

	writeCredReqSecret(credentialsRequest, outputDir, *userAssignedManagedIdentity.Properties.ClientID, *userAssignedManagedIdentity.Properties.TenantID, subscriptionID, region)
	return nil
}

// getRoleDefinitionByRoleName lists role definitions within the scope of the Azure subscription identified by subscriptionID
// and returns the armauthorization.RoleDefinition with a name matching the provided roleName.
//
// If multiple roles are found matching the roleName this will result in an error.
func getRoleDefinitionByRoleName(client *azureclients.AzureClientWrapper, roleName, subscriptionID string) (*armauthorization.RoleDefinition, error) {
	listRoles := client.RoleDefinitionsClient.NewListPager(
		"/subscriptions/"+subscriptionID,
		&armauthorization.RoleDefinitionsClientListOptions{
			Filter: to.Ptr(fmt.Sprintf("roleName eq '%v'", roleName)),
		},
	)
	roleDefinitions := make([]*armauthorization.RoleDefinition, 0)
	for listRoles.More() {
		pageResponse, err := listRoles.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		roleDefinitions = append(roleDefinitions, pageResponse.RoleDefinitionListResult.Value...)
	}
	switch len(roleDefinitions) {
	case 0:
		return nil, fmt.Errorf("no role found for name %q", roleName)
	case 1:
		return roleDefinitions[0], nil
	default:
		return nil, fmt.Errorf("found %q role definitions for %q, expected one", len(roleDefinitions), roleName)
	}
}

// getRoleDefinitionByID gets the role definition identified by roleDefinitionID.
func getRoleDefinitionByID(client *azureclients.AzureClientWrapper, roleDefinitionID string) (*armauthorization.RoleDefinition, error) {
	roleDefinitionGetResp, err := client.RoleDefinitionsClient.GetByID(
		context.Background(),
		roleDefinitionID,
		&armauthorization.RoleDefinitionsClientGetByIDOptions{},
	)
	if err != nil {
		return nil, err
	}
	return &roleDefinitionGetResp.RoleDefinition, nil
}

// createRoleAssignment assigns the Azure role with roleName/roleID to the managed identity identified by
// managedIdentityPrincipalID within the provided scope.
//
// Scope is a string such as /subscriptions/<subscriptionID> which represents anything within the subscription.
// This scope can be restricted within a resourceGroup such as /subscriptions/<subscriptionID>/resourceGroups/<resourceGroupName>.
func createRoleAssignment(client *azureclients.AzureClientWrapper, managedIdentityPrincipalID, roleID, roleName, scope, subscriptionID string) (*armauthorization.RoleAssignment, error) {
	// Create a unique name for the role assignment
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID for user-assigned identity role assignment: %v", err)
	}
	roleAssignmentName := uuid.String()

	var rawResponse *http.Response
	// Role assignment can fail due to a replication delay after creating the user-assigned managed identity
	// Try up to 6 times with a 10 second delay between each attempt, up to 1 minute.
	for i := 0; i < 6; i++ {
		ctxWithResp := runtime.WithCaptureResponse(context.Background(), &rawResponse)
		roleAssignmentCreateResponse, err := client.RoleAssignmentClient.Create(
			ctxWithResp,
			scope,
			roleAssignmentName,
			armauthorization.RoleAssignmentCreateParameters{
				Properties: &armauthorization.RoleAssignmentProperties{
					PrincipalID:      to.Ptr(managedIdentityPrincipalID),
					RoleDefinitionID: to.Ptr(roleID),
				},
			},
			&armauthorization.RoleAssignmentsClientCreateOptions{},
		)
		if err != nil {
			var respErr *azcore.ResponseError
			if errors.As(err, &respErr) {
				if respErr.ErrorCode == "PrincipalNotFound" {
					// The identity ccoctl just created can't be found yet due to a replication delay so we need to retry.
					if i == 5 {
						log.Fatal("Timed out assigning role to user-assigned managed identity, this is most likely due to a replication delay following creation of the user-assigned managed identity, please retry")
						break
					} else {
						log.Printf("Unable to assign role to user-assigned managed identity, retrying...")
						time.Sleep(10 * time.Second)
						continue
					}
				} else if respErr.ErrorCode == "RoleAssignmentExists" {
					log.Printf("Found existing role assignment %s for user-assigned managed identity with principal ID %s at scope %s", roleName, managedIdentityPrincipalID, scope)
					break
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			log.Printf("Created role assignment for role %s with user-assigned managed identity principal ID %s at scope %s", roleName, managedIdentityPrincipalID, scope)
			return &roleAssignmentCreateResponse.RoleAssignment, nil
		}
	}
	return nil, nil
}

// deleteRoleAssignment deletes the Azure role assignment with roleID and scope from the managed identity
// identified by managedIdentityPrincipalID
//
// Scope is a string such as /subscriptions/<subscriptionID> which represents anything within the subscription.
// This scope can be restricted within a resourceGroup such as /subscriptions/<subscriptionID>/resourceGroups/<resourceGroupName>.
func deleteRoleAssignment(client *azureclients.AzureClientWrapper, managedIdentityPrincipalID, roleID, roleName, scope, subscriptionID string) error {
	_, err := client.RoleAssignmentClient.Delete(
		context.Background(),
		scope,
		roleID,
		&armauthorization.RoleAssignmentsClientDeleteOptions{},
	)
	if err != nil {
		return err
	}
	log.Printf("Deleted role assignment for role %s with user-assigned managed identity with principal ID %s at scope %s", roleName, managedIdentityPrincipalID, scope)
	return nil
}

// ensureUserAssignedManagedIdentity ensures that a user-assigned managed identity with managedIdentityName exists
// within the provided resourceGroup
//
// resourceTags will be updated to match those provided to ensureUserAssignedManagedIdentity if found to be different on the existing user-assigned managed identity.
func ensureUserAssignedManagedIdentity(client *azureclients.AzureClientWrapper, managedIdentityName, resourceGroupName, region string, resourceTags map[string]string) (*armmsi.Identity, error) {
	// Check if user-assigned managed identity exists
	needToCreateUserAssignedManagedIdentity := false
	var rawResponse *http.Response
	ctxWithResp := runtime.WithCaptureResponse(context.Background(), &rawResponse)
	getUserAssignedManagedIdentityResp, err := client.UserAssignedIdentitiesClient.Get(
		ctxWithResp,
		resourceGroupName,
		managedIdentityName,
		&armmsi.UserAssignedIdentitiesClientGetOptions{})
	if err != nil {
		var respErr *azcore.ResponseError
		if errors.As(err, &respErr) {
			switch respErr.ErrorCode {
			case "ResourceNotFound":
				// User-assigned managed identity wasn't found and will need to be created
				needToCreateUserAssignedManagedIdentity = true
			default:
				return nil, errors.Wrapf(err, "unable to get user-assigned managed identity")
			}
		} else {
			return nil, err
		}
	}

	mergedResourceTags, needToUpdateUserAssignedManagedIdentity := mergeResourceTags(resourceTags, getUserAssignedManagedIdentityResp.Tags)

	// Found and validated existing user-assigned managed identity
	if !needToCreateUserAssignedManagedIdentity && !needToUpdateUserAssignedManagedIdentity {
		log.Printf("Found existing user-assigned managed identity %s", *getUserAssignedManagedIdentityResp.Identity.ID)
		return &getUserAssignedManagedIdentityResp.Identity, nil
	}

	identityParameters := armmsi.Identity{
		Location: to.Ptr(region),
		Tags:     mergedResourceTags,
	}

	userAssignedManagedIdentity, err := client.UserAssignedIdentitiesClient.CreateOrUpdate(
		context.Background(),
		resourceGroupName,
		managedIdentityName,
		identityParameters,
		&armmsi.UserAssignedIdentitiesClientCreateOrUpdateOptions{},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user-assigned managed identity")
	}
	verb := "Updated"
	if needToCreateUserAssignedManagedIdentity {
		verb = "Created"
	}
	log.Printf("%s user-assigned managed identity %s", verb, *userAssignedManagedIdentity.ID)
	return &userAssignedManagedIdentity.Identity, nil
}

// createFederatedIdentityCredential creates an Azure federated identity credential within the user-assigned managed
// identity identified by managedIdentityName.
//
// Federated identity credentials are limited to a specific kubernetes service account by providing the service account's
// name and namespace. The issuerURL of the OIDC endpoint hosting OIDC discovery and JWKS (public key information) documents
// must also be known to establish trust from a token signed by the OIDC endpoint's matching private key.
func createFederatedIdentityCredential(client *azureclients.AzureClientWrapper, managedIdentityName, issuerURL, serviceAccountNamespace, serviceAccountName, resourceGroupName string) error {
	federatedIdentityCredentialParameters := armmsi.FederatedIdentityCredential{
		Properties: &armmsi.FederatedIdentityCredentialProperties{
			Audiences: []*string{
				to.Ptr("openshift"),
			},
			Issuer:  to.Ptr(issuerURL),
			Subject: to.Ptr(fmt.Sprintf("system:serviceaccount:%s:%s", serviceAccountNamespace, serviceAccountName)),
		},
	}

	var (
		rawResponse                             *http.Response
		needToCreateFederatedIdentityCredential bool
	)
	ctxWithResp := runtime.WithCaptureResponse(context.Background(), &rawResponse)
	federatedIdentityCredentialGetResp, err := client.FederatedIdentityCredentialsClient.Get(
		ctxWithResp,
		resourceGroupName,
		managedIdentityName,
		serviceAccountName,
		&armmsi.FederatedIdentityCredentialsClientGetOptions{},
	)
	if err != nil {
		var respErr *azcore.ResponseError
		if errors.As(err, &respErr) {
			switch respErr.ErrorCode {
			case "ResourceNotFound":
				// Federated identity credential wasn't found and will need to be created
				needToCreateFederatedIdentityCredential = true
			default:
				return errors.Wrapf(err, "unable to get federated identity credential")
			}
		} else {
			return err
		}
	}

	if !needToCreateFederatedIdentityCredential {
		log.Printf("Found existing federated identity credential %s", *federatedIdentityCredentialGetResp.FederatedIdentityCredential.ID)
		return nil
	}

	federatedIdentityCredential, err := client.FederatedIdentityCredentialsClient.CreateOrUpdate(
		context.Background(),
		resourceGroupName,
		managedIdentityName,
		serviceAccountName,
		federatedIdentityCredentialParameters,
		&armmsi.FederatedIdentityCredentialsClientCreateOrUpdateOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to create federated identity credential")
	}
	log.Printf("Created federated identity credential %s", *federatedIdentityCredential.ID)
	return nil
}

// writeCredReqSecret writes a secret file within the manifests directory (outputDir/manifests/)
// containing user-assigned managed identity details.
func writeCredReqSecret(cr *credreqv1.CredentialsRequest, outputDir, clientID, tenantID, subscriptionID, region string) error {
	manifestsDir := filepath.Join(outputDir, provisioning.ManifestsDirName)
	fileName := fmt.Sprintf("%s-%s-credentials.yaml", cr.Spec.SecretRef.Namespace, cr.Spec.SecretRef.Name)
	filePath := filepath.Join(manifestsDir, fileName)
	fileData := fmt.Sprintf(secretManifestTemplate, clientID, tenantID, region, subscriptionID, provisioning.OidcTokenPath, cr.Spec.SecretRef.Name, cr.Spec.SecretRef.Namespace)

	// clientID would be an empty string if ccoctl was in --dry-run mode
	// so lets make sure we have an invalid Secret until the user
	// has populated the secret manually
	if clientID == "" && tenantID == "" {
		fileData = fileData + "\nPOPULATE CLIENT ID AND TENANT ID AND DELETE THIS LINE"
	}

	if err := os.WriteFile(filePath, []byte(fileData), 0600); err != nil {
		return errors.Wrapf(err, "failed to save secret file at path %s", filePath)
	}
	log.Printf("Saved credentials configuration to: %s", filePath)

	return nil
}

// createManagedIdentities creates user-assigned managed identities for each CredentialsRequest found within the creqReqDir.
//
// User-assigned managed identities are created within the resource group identified by oidcResourceGroupName.
//
// Roles listed within the CredentialsRequest (spec.providerSpec.roleBindings) will be assigned to created user-assigned
// managed identities and role assignment will be scoped to the resource group identified by installationResourceGroupName
// by default.
//
// Role assignment for the specific CredentialsRequest named "openshift-ingress-azure" (ingressCredentialsRequestName) will be
// additionally scoped within the resource group identified by dnsZoneResourceGroupName.
//
// Kubernetes secrets containing the user-assigned managed identity's clientID will be generated and written to the outputDir.
func createManagedIdentities(client *azureclients.AzureClientWrapper, credReqDir, name, oidcResourceGroupName, subscriptionID, region, issuerURL, outputDir, installationResourceGroupName, dnsZoneResourceGroupName string, resourceTags map[string]string, enableTechPreview, dryRun bool) error {
	// Add CCO's "owned" tag to resource tags map
	resourceTags[fmt.Sprintf("%s_%s", ownedAzureResourceTagKeyPrefix, name)] = ownedAzureResourceTagValue

	// Ensure the installation resource group exists
	if !dryRun {
		err := ensureResourceGroup(client, installationResourceGroupName, region, resourceTags)
		if err != nil {
			return errors.Wrap(err, "failed to ensure resource group")
		}
		log.Printf("Cluster installation resource group name is %s. This resource group MUST be configured as the resource group used for cluster installation.", installationResourceGroupName)
	}

	// Process directory containing CredentialsRequests object manifests into list of CredentialsRequests objects
	credentialsRequests, err := provisioning.GetListOfCredentialsRequests(credReqDir, enableTechPreview)
	if err != nil {
		return errors.Wrap(err, "failed to process files containing CredentialsRequests")
	}

	// Create user-assigned managed identities for each CredentialsRequest
	for _, credentialsRequest := range credentialsRequests {
		// Scope user-assigned managed identity within the installationResourceGroupName
		scopingResourceGroupNames := []string{installationResourceGroupName}
		// Additionally scope the ingress CredentialsRequest within the dnsZoneResourceGroupName
		if credentialsRequest.Name == ingressCredentialRequestName {
			scopingResourceGroupNames = append(scopingResourceGroupNames, dnsZoneResourceGroupName)
		}
		err := createManagedIdentity(client, name, oidcResourceGroupName, subscriptionID, region, issuerURL, outputDir, scopingResourceGroupNames, resourceTags, credentialsRequest, dryRun)
		if err != nil {
			return err
		}
	}

	return nil
}

func createManagedIdentitiesCmd(cmd *cobra.Command, args []string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	azureClientWrapper, err := azureclients.NewAzureClientWrapper(CreateManagedIdentitiesOpts.SubscriptionID, cred, &policy.ClientOptions{}, false)
	if err != nil {
		log.Fatalf("Failed to create Azure client: %s", err)
	}

	if CreateManagedIdentitiesOpts.OIDCResourceGroupName == "" {
		CreateManagedIdentitiesOpts.OIDCResourceGroupName = CreateManagedIdentitiesOpts.Name + oidcResourceGroupSuffix
		log.Printf("No --oidc-resource-group-name provided, defaulting OIDC resource group name to %s", CreateManagedIdentitiesOpts.OIDCResourceGroupName)
	}

	if CreateManagedIdentitiesOpts.InstallationResourceGroupName == "" {
		CreateManagedIdentitiesOpts.InstallationResourceGroupName = CreateManagedIdentitiesOpts.Name
		log.Printf("No --installation-resource-group-name provided, defaulting installation resource group name to %s", CreateManagedIdentitiesOpts.InstallationResourceGroupName)
	}

	err = createManagedIdentities(
		azureClientWrapper,
		CreateManagedIdentitiesOpts.CredRequestDir,
		CreateManagedIdentitiesOpts.Name,
		CreateManagedIdentitiesOpts.OIDCResourceGroupName,
		CreateManagedIdentitiesOpts.SubscriptionID,
		CreateManagedIdentitiesOpts.Region,
		CreateManagedIdentitiesOpts.IssuerURL,
		CreateManagedIdentitiesOpts.OutputDir,
		CreateManagedIdentitiesOpts.InstallationResourceGroupName,
		CreateManagedIdentitiesOpts.DNSZoneResourceGroupName,
		CreateManagedIdentitiesOpts.UserTags,
		CreateManagedIdentitiesOpts.EnableTechPreview,
		CreateManagedIdentitiesOpts.DryRun)
	if err != nil {
		log.Fatal(err)
	}
}

// initEnvForCreateManagedIdentitiesCmd ensures that the output directory specified by --output-dir exists
func initEnvForCreateManagedIdentitiesCmd(cmd *cobra.Command, args []string) {
	if CreateManagedIdentitiesOpts.OutputDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current directory: %s", err)
		}
		CreateManagedIdentitiesOpts.OutputDir = pwd
		log.Printf("No --output-dir provided, defaulting output directory to the current working directory %s", CreateManagedIdentitiesOpts.OutputDir)
	}

	outputDirPath, err := filepath.Abs(CreateManagedIdentitiesOpts.OutputDir)
	if err != nil {
		log.Fatalf("Failed to resolve full path of the output directory %s", err)
	}

	// Create output dir if it doesn't exist
	err = provisioning.EnsureDir(outputDirPath)
	if err != nil {
		log.Fatalf("Failed to create target directory at path %s", outputDirPath)
	}

	// Create manifests dir within the output dir if it doesn't exist
	manifestsDir := filepath.Join(outputDirPath, provisioning.ManifestsDirName)
	err = provisioning.EnsureDir(manifestsDir)
	if err != nil {
		log.Fatalf("Failed to create manifests directory at path %s", manifestsDir)
	}
}

// NewCreateManagedIdentitiesCmd provides the "create-managed-identities" subcommand
func NewCreateManagedIdentitiesCmd() *cobra.Command {
	createManagedIdentitiesCmd := &cobra.Command{
		Use:              "create-managed-identities --name NAME --region REGION --credentials-requests-dir CRED_REQ_DIR \\ \n\t--installation-resource-group-name RESOURCE_GROUP_NAME --dnszone-resource-group-name RESOURCE_GROUP_NAME \\ \n\t--issuer-url ISSUER_URL --subscription-id SUBSCRIPTION_ID",
		Short:            "Create Azure Managed Identities",
		Run:              createManagedIdentitiesCmd,
		PersistentPreRun: initEnvForCreateManagedIdentitiesCmd,
	}

	// Required
	createManagedIdentitiesCmd.PersistentFlags().StringVar(
		&CreateManagedIdentitiesOpts.Name,
		"name",
		"",
		"User-defined name for all created Azure resources. This user-defined name can be separate from the cluster's infra-id. "+
			fmt.Sprintf("Azure resources created by ccoctl will be tagged with '%s_NAME = %s'", ownedAzureResourceTagKeyPrefix, ownedAzureResourceTagValue),
	)
	createManagedIdentitiesCmd.MarkPersistentFlagRequired("name")
	createManagedIdentitiesCmd.PersistentFlags().StringVar(&CreateManagedIdentitiesOpts.Region, "region", "", "Azure region in which to create user-assigned managed identities")
	createManagedIdentitiesCmd.MarkPersistentFlagRequired("region")
	createManagedIdentitiesCmd.PersistentFlags().StringVar(&CreateManagedIdentitiesOpts.CredRequestDir, "credentials-requests-dir", "", "Directory containing Azure CredentialsRequests files used to create user-assigned managed identities (can be created by running 'oc adm release extract --credentials-requests --cloud=azure' against an OpenShift release image)")
	createManagedIdentitiesCmd.MarkPersistentFlagRequired("credentials-requests-dir")
	createManagedIdentitiesCmd.PersistentFlags().StringVar(&CreateManagedIdentitiesOpts.DNSZoneResourceGroupName, "dnszone-resource-group-name", "", "The existing Azure resource group which contains the DNS zone that will be used for the cluster's base domain. The cluster ingress operator will be scoped to allow management of DNS records in the DNS Zone resource group.")
	createManagedIdentitiesCmd.MarkPersistentFlagRequired("dnszone-resource-group-name")
	createManagedIdentitiesCmd.PersistentFlags().StringVar(
		&CreateManagedIdentitiesOpts.InstallationResourceGroupName,
		"installation-resource-group-name",
		"",
		"The Azure resource group which will be used for future cluster installation. "+
			"Managed identities will be scoped such that they can manage resources in this resource group. "+
			"The OpenShift installer requires that the resource group provided for installation resources be initially empty so this resource group must "+
			"contain no resources if the resource group was previously created. "+
			"A resource group will be created (with name derived from the --name parameter) if an installation-resource-group-name parameter was not provided. "+
			"Note that this resource group must be provided as the installation resource group when installing the OpenShift cluster",
	)
	createManagedIdentitiesCmd.PersistentFlags().StringVar(&CreateManagedIdentitiesOpts.SubscriptionID, "subscription-id", "", "Azure Subscription ID within which to create and scope the access of managed identities")
	createManagedIdentitiesCmd.MarkPersistentFlagRequired("subscription-id")
	createManagedIdentitiesCmd.PersistentFlags().StringVar(&CreateManagedIdentitiesOpts.IssuerURL, "issuer-url", "", "OIDC Issuer URL (the OIDC Issuer can be created with the 'create-oidc-issuer' sub-command)")
	createManagedIdentitiesCmd.MarkPersistentFlagRequired("issuer-url")

	// Optional
	createManagedIdentitiesCmd.PersistentFlags().StringVar(
		&CreateManagedIdentitiesOpts.OIDCResourceGroupName,
		"oidc-resource-group-name",
		"",
		"The Azure resource group resource group in which to create user-assigned managed identities (can be created with the 'create-oidc-issuer' sub-command). "+
			"A resource group will be created with a name derived from the --name parameter if an --oidc-resource-group-name parameter was not provided.",
	)
	createManagedIdentitiesCmd.PersistentFlags().BoolVar(&CreateManagedIdentitiesOpts.DryRun, "dry-run", false, "Skip creating objects and just save what would have been created into files")
	createManagedIdentitiesCmd.PersistentFlags().StringVar(&CreateManagedIdentitiesOpts.OutputDir, "output-dir", "", "Directory to place generated files. Defaults to the current directory.")
	createManagedIdentitiesCmd.PersistentFlags().StringToStringVar(&CreateManagedIdentitiesOpts.UserTags, "user-tags", map[string]string{}, "User tags to be applied to Azure resources, multiple tags may be specified comma-separated for example: --user-tags key1=value1,key2=value2")
	createManagedIdentitiesCmd.PersistentFlags().BoolVar(&CreateManagedIdentitiesOpts.EnableTechPreview, "enable-tech-preview", false, "Opt into processing CredentialsRequests annotated with TechPreviewNoUpgrade")

	return createManagedIdentitiesCmd
}
