// Code generated by MockGen. DO NOT EDIT.
// Source: ./clients.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armauthorization "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
	armmsi "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	armresources "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	blockblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	gomock "github.com/golang/mock/gomock"
	models "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// MockAppClient is a mock of AppClient interface.
type MockAppClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppClientMockRecorder
}

// MockAppClientMockRecorder is the mock recorder for MockAppClient.
type MockAppClientMockRecorder struct {
	mock *MockAppClient
}

// NewMockAppClient creates a new mock instance.
func NewMockAppClient(ctrl *gomock.Controller) *MockAppClient {
	mock := &MockAppClient{ctrl: ctrl}
	mock.recorder = &MockAppClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppClient) EXPECT() *MockAppClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockAppClient) Delete(ctx context.Context, applicationObjectID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, applicationObjectID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAppClientMockRecorder) Delete(ctx, applicationObjectID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAppClient)(nil).Delete), ctx, applicationObjectID)
}

// List mocks base method.
func (m *MockAppClient) List(ctx context.Context, filter string) ([]models.Applicationable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, filter)
	ret0, _ := ret[0].([]models.Applicationable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAppClientMockRecorder) List(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAppClient)(nil).List), ctx, filter)
}

// MockResourceGroupsClient is a mock of ResourceGroupsClient interface.
type MockResourceGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourceGroupsClientMockRecorder
}

// MockResourceGroupsClientMockRecorder is the mock recorder for MockResourceGroupsClient.
type MockResourceGroupsClientMockRecorder struct {
	mock *MockResourceGroupsClient
}

// NewMockResourceGroupsClient creates a new mock instance.
func NewMockResourceGroupsClient(ctrl *gomock.Controller) *MockResourceGroupsClient {
	mock := &MockResourceGroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourceGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceGroupsClient) EXPECT() *MockResourceGroupsClientMockRecorder {
	return m.recorder
}

// BeginDelete mocks base method.
func (m *MockResourceGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, options *armresources.ResourceGroupsClientBeginDeleteOptions) (*runtime.Poller[armresources.ResourceGroupsClientDeleteResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginDelete", ctx, resourceGroupName, options)
	ret0, _ := ret[0].(*runtime.Poller[armresources.ResourceGroupsClientDeleteResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginDelete indicates an expected call of BeginDelete.
func (mr *MockResourceGroupsClientMockRecorder) BeginDelete(ctx, resourceGroupName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginDelete", reflect.TypeOf((*MockResourceGroupsClient)(nil).BeginDelete), ctx, resourceGroupName, options)
}

// CreateOrUpdate mocks base method.
func (m *MockResourceGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters armresources.ResourceGroup, options *armresources.ResourceGroupsClientCreateOrUpdateOptions) (armresources.ResourceGroupsClientCreateOrUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, parameters, options)
	ret0, _ := ret[0].(armresources.ResourceGroupsClientCreateOrUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockResourceGroupsClientMockRecorder) CreateOrUpdate(ctx, resourceGroupName, parameters, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockResourceGroupsClient)(nil).CreateOrUpdate), ctx, resourceGroupName, parameters, options)
}

// Get mocks base method.
func (m *MockResourceGroupsClient) Get(ctx context.Context, resourceGroupName string, options *armresources.ResourceGroupsClientGetOptions) (armresources.ResourceGroupsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, options)
	ret0, _ := ret[0].(armresources.ResourceGroupsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockResourceGroupsClientMockRecorder) Get(ctx, resourceGroupName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceGroupsClient)(nil).Get), ctx, resourceGroupName, options)
}

// MockAccountsClient is a mock of AccountsClient interface.
type MockAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsClientMockRecorder
}

// MockAccountsClientMockRecorder is the mock recorder for MockAccountsClient.
type MockAccountsClientMockRecorder struct {
	mock *MockAccountsClient
}

// NewMockAccountsClient creates a new mock instance.
func NewMockAccountsClient(ctrl *gomock.Controller) *MockAccountsClient {
	mock := &MockAccountsClient{ctrl: ctrl}
	mock.recorder = &MockAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsClient) EXPECT() *MockAccountsClientMockRecorder {
	return m.recorder
}

// BeginCreate mocks base method.
func (m *MockAccountsClient) BeginCreate(ctx context.Context, resourceGroupName, accountName string, parameters armstorage.AccountCreateParameters, options *armstorage.AccountsClientBeginCreateOptions) (*runtime.Poller[armstorage.AccountsClientCreateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginCreate", ctx, resourceGroupName, accountName, parameters, options)
	ret0, _ := ret[0].(*runtime.Poller[armstorage.AccountsClientCreateResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginCreate indicates an expected call of BeginCreate.
func (mr *MockAccountsClientMockRecorder) BeginCreate(ctx, resourceGroupName, accountName, parameters, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginCreate", reflect.TypeOf((*MockAccountsClient)(nil).BeginCreate), ctx, resourceGroupName, accountName, parameters, options)
}

// Delete mocks base method.
func (m *MockAccountsClient) Delete(ctx context.Context, resourceGroupName, accountName string, options *armstorage.AccountsClientDeleteOptions) (armstorage.AccountsClientDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, accountName, options)
	ret0, _ := ret[0].(armstorage.AccountsClientDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountsClientMockRecorder) Delete(ctx, resourceGroupName, accountName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountsClient)(nil).Delete), ctx, resourceGroupName, accountName, options)
}

// ListKeys mocks base method.
func (m *MockAccountsClient) ListKeys(ctx context.Context, resourceGroupName, accountName string, options *armstorage.AccountsClientListKeysOptions) (armstorage.AccountsClientListKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", ctx, resourceGroupName, accountName, options)
	ret0, _ := ret[0].(armstorage.AccountsClientListKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockAccountsClientMockRecorder) ListKeys(ctx, resourceGroupName, accountName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockAccountsClient)(nil).ListKeys), ctx, resourceGroupName, accountName, options)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockAccountsClient) NewListByResourceGroupPager(resourceGroupName string, options *armstorage.AccountsClientListByResourceGroupOptions) *runtime.Pager[armstorage.AccountsClientListByResourceGroupResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", resourceGroupName, options)
	ret0, _ := ret[0].(*runtime.Pager[armstorage.AccountsClientListByResourceGroupResponse])
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockAccountsClientMockRecorder) NewListByResourceGroupPager(resourceGroupName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockAccountsClient)(nil).NewListByResourceGroupPager), resourceGroupName, options)
}

// NewListPager mocks base method.
func (m *MockAccountsClient) NewListPager(options *armstorage.AccountsClientListOptions) *runtime.Pager[armstorage.AccountsClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", options)
	ret0, _ := ret[0].(*runtime.Pager[armstorage.AccountsClientListResponse])
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockAccountsClientMockRecorder) NewListPager(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockAccountsClient)(nil).NewListPager), options)
}

// Update mocks base method.
func (m *MockAccountsClient) Update(ctx context.Context, resourceGroupName, accountName string, parameters armstorage.AccountUpdateParameters, options *armstorage.AccountsClientUpdateOptions) (armstorage.AccountsClientUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, resourceGroupName, accountName, parameters, options)
	ret0, _ := ret[0].(armstorage.AccountsClientUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountsClientMockRecorder) Update(ctx, resourceGroupName, accountName, parameters, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountsClient)(nil).Update), ctx, resourceGroupName, accountName, parameters, options)
}

// MockBlobContainersClient is a mock of BlobContainersClient interface.
type MockBlobContainersClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlobContainersClientMockRecorder
}

// MockBlobContainersClientMockRecorder is the mock recorder for MockBlobContainersClient.
type MockBlobContainersClientMockRecorder struct {
	mock *MockBlobContainersClient
}

// NewMockBlobContainersClient creates a new mock instance.
func NewMockBlobContainersClient(ctrl *gomock.Controller) *MockBlobContainersClient {
	mock := &MockBlobContainersClient{ctrl: ctrl}
	mock.recorder = &MockBlobContainersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlobContainersClient) EXPECT() *MockBlobContainersClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBlobContainersClient) Create(ctx context.Context, resourceGroupName, accountName, containerName string, blobContainer armstorage.BlobContainer, options *armstorage.BlobContainersClientCreateOptions) (armstorage.BlobContainersClientCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, resourceGroupName, accountName, containerName, blobContainer, options)
	ret0, _ := ret[0].(armstorage.BlobContainersClientCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBlobContainersClientMockRecorder) Create(ctx, resourceGroupName, accountName, containerName, blobContainer, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlobContainersClient)(nil).Create), ctx, resourceGroupName, accountName, containerName, blobContainer, options)
}

// Get mocks base method.
func (m *MockBlobContainersClient) Get(ctx context.Context, resourceGroupName, accountName, containerName string, options *armstorage.BlobContainersClientGetOptions) (armstorage.BlobContainersClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, accountName, containerName, options)
	ret0, _ := ret[0].(armstorage.BlobContainersClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBlobContainersClientMockRecorder) Get(ctx, resourceGroupName, accountName, containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBlobContainersClient)(nil).Get), ctx, resourceGroupName, accountName, containerName, options)
}

// MockAZBlobClient is a mock of AZBlobClient interface.
type MockAZBlobClient struct {
	ctrl     *gomock.Controller
	recorder *MockAZBlobClientMockRecorder
}

// MockAZBlobClientMockRecorder is the mock recorder for MockAZBlobClient.
type MockAZBlobClientMockRecorder struct {
	mock *MockAZBlobClient
}

// NewMockAZBlobClient creates a new mock instance.
func NewMockAZBlobClient(ctrl *gomock.Controller) *MockAZBlobClient {
	mock := &MockAZBlobClient{ctrl: ctrl}
	mock.recorder = &MockAZBlobClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAZBlobClient) EXPECT() *MockAZBlobClientMockRecorder {
	return m.recorder
}

// UploadBuffer mocks base method.
func (m *MockAZBlobClient) UploadBuffer(ctx context.Context, containerName, blobName string, buffer []byte, o *blockblob.UploadBufferOptions) (blockblob.UploadBufferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadBuffer", ctx, containerName, blobName, buffer, o)
	ret0, _ := ret[0].(blockblob.UploadBufferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadBuffer indicates an expected call of UploadBuffer.
func (mr *MockAZBlobClientMockRecorder) UploadBuffer(ctx, containerName, blobName, buffer, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadBuffer", reflect.TypeOf((*MockAZBlobClient)(nil).UploadBuffer), ctx, containerName, blobName, buffer, o)
}

// MockUserAssignedIdentitiesClient is a mock of UserAssignedIdentitiesClient interface.
type MockUserAssignedIdentitiesClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserAssignedIdentitiesClientMockRecorder
}

// MockUserAssignedIdentitiesClientMockRecorder is the mock recorder for MockUserAssignedIdentitiesClient.
type MockUserAssignedIdentitiesClientMockRecorder struct {
	mock *MockUserAssignedIdentitiesClient
}

// NewMockUserAssignedIdentitiesClient creates a new mock instance.
func NewMockUserAssignedIdentitiesClient(ctrl *gomock.Controller) *MockUserAssignedIdentitiesClient {
	mock := &MockUserAssignedIdentitiesClient{ctrl: ctrl}
	mock.recorder = &MockUserAssignedIdentitiesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAssignedIdentitiesClient) EXPECT() *MockUserAssignedIdentitiesClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockUserAssignedIdentitiesClient) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName string, parameters armmsi.Identity, options *armmsi.UserAssignedIdentitiesClientCreateOrUpdateOptions) (armmsi.UserAssignedIdentitiesClientCreateOrUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, parameters, options)
	ret0, _ := ret[0].(armmsi.UserAssignedIdentitiesClientCreateOrUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockUserAssignedIdentitiesClientMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, parameters, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockUserAssignedIdentitiesClient)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, parameters, options)
}

// Delete mocks base method.
func (m *MockUserAssignedIdentitiesClient) Delete(ctx context.Context, resourceGroupName, resourceName string, options *armmsi.UserAssignedIdentitiesClientDeleteOptions) (armmsi.UserAssignedIdentitiesClientDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, resourceName, options)
	ret0, _ := ret[0].(armmsi.UserAssignedIdentitiesClientDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserAssignedIdentitiesClientMockRecorder) Delete(ctx, resourceGroupName, resourceName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserAssignedIdentitiesClient)(nil).Delete), ctx, resourceGroupName, resourceName, options)
}

// Get mocks base method.
func (m *MockUserAssignedIdentitiesClient) Get(ctx context.Context, resourceGroupName, resourceName string, options *armmsi.UserAssignedIdentitiesClientGetOptions) (armmsi.UserAssignedIdentitiesClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, options)
	ret0, _ := ret[0].(armmsi.UserAssignedIdentitiesClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserAssignedIdentitiesClientMockRecorder) Get(ctx, resourceGroupName, resourceName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserAssignedIdentitiesClient)(nil).Get), ctx, resourceGroupName, resourceName, options)
}

// NewListByResourceGroupPager mocks base method.
func (m *MockUserAssignedIdentitiesClient) NewListByResourceGroupPager(resourceGroupName string, options *armmsi.UserAssignedIdentitiesClientListByResourceGroupOptions) *runtime.Pager[armmsi.UserAssignedIdentitiesClientListByResourceGroupResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListByResourceGroupPager", resourceGroupName, options)
	ret0, _ := ret[0].(*runtime.Pager[armmsi.UserAssignedIdentitiesClientListByResourceGroupResponse])
	return ret0
}

// NewListByResourceGroupPager indicates an expected call of NewListByResourceGroupPager.
func (mr *MockUserAssignedIdentitiesClientMockRecorder) NewListByResourceGroupPager(resourceGroupName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListByResourceGroupPager", reflect.TypeOf((*MockUserAssignedIdentitiesClient)(nil).NewListByResourceGroupPager), resourceGroupName, options)
}

// MockRoleDefinitionsClient is a mock of RoleDefinitionsClient interface.
type MockRoleDefinitionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleDefinitionsClientMockRecorder
}

// MockRoleDefinitionsClientMockRecorder is the mock recorder for MockRoleDefinitionsClient.
type MockRoleDefinitionsClientMockRecorder struct {
	mock *MockRoleDefinitionsClient
}

// NewMockRoleDefinitionsClient creates a new mock instance.
func NewMockRoleDefinitionsClient(ctrl *gomock.Controller) *MockRoleDefinitionsClient {
	mock := &MockRoleDefinitionsClient{ctrl: ctrl}
	mock.recorder = &MockRoleDefinitionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleDefinitionsClient) EXPECT() *MockRoleDefinitionsClientMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockRoleDefinitionsClient) GetByID(ctx context.Context, roleDefinitionID string, options *armauthorization.RoleDefinitionsClientGetByIDOptions) (armauthorization.RoleDefinitionsClientGetByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, roleDefinitionID, options)
	ret0, _ := ret[0].(armauthorization.RoleDefinitionsClientGetByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRoleDefinitionsClientMockRecorder) GetByID(ctx, roleDefinitionID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRoleDefinitionsClient)(nil).GetByID), ctx, roleDefinitionID, options)
}

// NewListPager mocks base method.
func (m *MockRoleDefinitionsClient) NewListPager(scope string, options *armauthorization.RoleDefinitionsClientListOptions) *runtime.Pager[armauthorization.RoleDefinitionsClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", scope, options)
	ret0, _ := ret[0].(*runtime.Pager[armauthorization.RoleDefinitionsClientListResponse])
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockRoleDefinitionsClientMockRecorder) NewListPager(scope, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockRoleDefinitionsClient)(nil).NewListPager), scope, options)
}

// MockRoleAssignmentsClient is a mock of RoleAssignmentsClient interface.
type MockRoleAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoleAssignmentsClientMockRecorder
}

// MockRoleAssignmentsClientMockRecorder is the mock recorder for MockRoleAssignmentsClient.
type MockRoleAssignmentsClientMockRecorder struct {
	mock *MockRoleAssignmentsClient
}

// NewMockRoleAssignmentsClient creates a new mock instance.
func NewMockRoleAssignmentsClient(ctrl *gomock.Controller) *MockRoleAssignmentsClient {
	mock := &MockRoleAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockRoleAssignmentsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleAssignmentsClient) EXPECT() *MockRoleAssignmentsClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoleAssignmentsClient) Create(ctx context.Context, scope, roleAssignmentName string, parameters armauthorization.RoleAssignmentCreateParameters, options *armauthorization.RoleAssignmentsClientCreateOptions) (armauthorization.RoleAssignmentsClientCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, scope, roleAssignmentName, parameters, options)
	ret0, _ := ret[0].(armauthorization.RoleAssignmentsClientCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleAssignmentsClientMockRecorder) Create(ctx, scope, roleAssignmentName, parameters, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).Create), ctx, scope, roleAssignmentName, parameters, options)
}

// Delete mocks base method.
func (m *MockRoleAssignmentsClient) Delete(ctx context.Context, scope, roleAssignmentName string, options *armauthorization.RoleAssignmentsClientDeleteOptions) (armauthorization.RoleAssignmentsClientDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, scope, roleAssignmentName, options)
	ret0, _ := ret[0].(armauthorization.RoleAssignmentsClientDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleAssignmentsClientMockRecorder) Delete(ctx, scope, roleAssignmentName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).Delete), ctx, scope, roleAssignmentName, options)
}

// NewListPager mocks base method.
func (m *MockRoleAssignmentsClient) NewListPager(options *armauthorization.RoleAssignmentsClientListOptions) *runtime.Pager[armauthorization.RoleAssignmentsClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", options)
	ret0, _ := ret[0].(*runtime.Pager[armauthorization.RoleAssignmentsClientListResponse])
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockRoleAssignmentsClientMockRecorder) NewListPager(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockRoleAssignmentsClient)(nil).NewListPager), options)
}

// MockFederatedIdentityCredentialsClient is a mock of FederatedIdentityCredentialsClient interface.
type MockFederatedIdentityCredentialsClient struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedIdentityCredentialsClientMockRecorder
}

// MockFederatedIdentityCredentialsClientMockRecorder is the mock recorder for MockFederatedIdentityCredentialsClient.
type MockFederatedIdentityCredentialsClientMockRecorder struct {
	mock *MockFederatedIdentityCredentialsClient
}

// NewMockFederatedIdentityCredentialsClient creates a new mock instance.
func NewMockFederatedIdentityCredentialsClient(ctrl *gomock.Controller) *MockFederatedIdentityCredentialsClient {
	mock := &MockFederatedIdentityCredentialsClient{ctrl: ctrl}
	mock.recorder = &MockFederatedIdentityCredentialsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedIdentityCredentialsClient) EXPECT() *MockFederatedIdentityCredentialsClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockFederatedIdentityCredentialsClient) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName, federatedIdentityCredentialResourceName string, parameters armmsi.FederatedIdentityCredential, options *armmsi.FederatedIdentityCredentialsClientCreateOrUpdateOptions) (armmsi.FederatedIdentityCredentialsClientCreateOrUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options)
	ret0, _ := ret[0].(armmsi.FederatedIdentityCredentialsClientCreateOrUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockFederatedIdentityCredentialsClientMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockFederatedIdentityCredentialsClient)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options)
}

// Get mocks base method.
func (m *MockFederatedIdentityCredentialsClient) Get(ctx context.Context, resourceGroupName, resourceName, federatedIdentityCredentialResourceName string, options *armmsi.FederatedIdentityCredentialsClientGetOptions) (armmsi.FederatedIdentityCredentialsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
	ret0, _ := ret[0].(armmsi.FederatedIdentityCredentialsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFederatedIdentityCredentialsClientMockRecorder) Get(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFederatedIdentityCredentialsClient)(nil).Get), ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
}

// MockMockablePoller is a mock of MockablePoller interface.
type MockMockablePoller[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockMockablePollerMockRecorder[T]
}

// MockMockablePollerMockRecorder is the mock recorder for MockMockablePoller.
type MockMockablePollerMockRecorder[T any] struct {
	mock *MockMockablePoller[T]
}

// NewMockMockablePoller creates a new mock instance.
func NewMockMockablePoller[T any](ctrl *gomock.Controller) *MockMockablePoller[T] {
	mock := &MockMockablePoller[T]{ctrl: ctrl}
	mock.recorder = &MockMockablePollerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMockablePoller[T]) EXPECT() *MockMockablePollerMockRecorder[T] {
	return m.recorder
}

// Done mocks base method.
func (m *MockMockablePoller[T]) Done() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockMockablePollerMockRecorder[T]) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockMockablePoller[T])(nil).Done))
}

// Poll mocks base method.
func (m *MockMockablePoller[T]) Poll(ctx context.Context) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Poll", ctx)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Poll indicates an expected call of Poll.
func (mr *MockMockablePollerMockRecorder[T]) Poll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Poll", reflect.TypeOf((*MockMockablePoller[T])(nil).Poll), ctx)
}

// PollUntilDone mocks base method.
func (m *MockMockablePoller[T]) PollUntilDone(ctx context.Context, options *runtime.PollUntilDoneOptions) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollUntilDone", ctx, options)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollUntilDone indicates an expected call of PollUntilDone.
func (mr *MockMockablePollerMockRecorder[T]) PollUntilDone(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollUntilDone", reflect.TypeOf((*MockMockablePoller[T])(nil).PollUntilDone), ctx, options)
}

// Result mocks base method.
func (m *MockMockablePoller[T]) Result(ctx context.Context) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result", ctx)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Result indicates an expected call of Result.
func (mr *MockMockablePollerMockRecorder[T]) Result(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockMockablePoller[T])(nil).Result), ctx)
}

// ResumeToken mocks base method.
func (m *MockMockablePoller[T]) ResumeToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResumeToken indicates an expected call of ResumeToken.
func (mr *MockMockablePollerMockRecorder[T]) ResumeToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeToken", reflect.TypeOf((*MockMockablePoller[T])(nil).ResumeToken))
}
