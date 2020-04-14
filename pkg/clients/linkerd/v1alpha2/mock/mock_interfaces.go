// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_v1alpha2 is a generated GoMock package.
package mock_v1alpha2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockServiceProfileClient is a mock of ServiceProfileClient interface.
type MockServiceProfileClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProfileClientMockRecorder
}

// MockServiceProfileClientMockRecorder is the mock recorder for MockServiceProfileClient.
type MockServiceProfileClientMockRecorder struct {
	mock *MockServiceProfileClient
}

// NewMockServiceProfileClient creates a new mock instance.
func NewMockServiceProfileClient(ctrl *gomock.Controller) *MockServiceProfileClient {
	mock := &MockServiceProfileClient{ctrl: ctrl}
	mock.recorder = &MockServiceProfileClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceProfileClient) EXPECT() *MockServiceProfileClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockServiceProfileClient) Get(ctx context.Context, key client.ObjectKey) (*v1alpha2.ServiceProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.ServiceProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceProfileClientMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceProfileClient)(nil).Get), ctx, key)
}

// List mocks base method.
func (m *MockServiceProfileClient) List(ctx context.Context, options ...client.ListOption) (*v1alpha2.ServiceProfileList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1alpha2.ServiceProfileList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceProfileClientMockRecorder) List(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceProfileClient)(nil).List), varargs...)
}

// CreateMeshWorkload mocks base method.
func (m *MockServiceProfileClient) Create(ctx context.Context, serviceProfile *v1alpha2.ServiceProfile, options ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, serviceProfile}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMeshWorkload", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMeshWorkload indicates an expected call of CreateMeshWorkload.
func (mr *MockServiceProfileClientMockRecorder) Create(ctx, serviceProfile interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, serviceProfile}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeshWorkload", reflect.TypeOf((*MockServiceProfileClient)(nil).Create), varargs...)
}

// UpdateService mocks base method.
func (m *MockServiceProfileClient) Update(ctx context.Context, serviceProfile *v1alpha2.ServiceProfile, options ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, serviceProfile}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateService", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockServiceProfileClientMockRecorder) Update(ctx, serviceProfile interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, serviceProfile}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockServiceProfileClient)(nil).Update), varargs...)
}

// UpsertSpec mocks base method.
func (m *MockServiceProfileClient) UpsertSpec(ctx context.Context, serviceProfile *v1alpha2.ServiceProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSpec", ctx, serviceProfile)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSpec indicates an expected call of UpsertSpec.
func (mr *MockServiceProfileClientMockRecorder) UpsertSpec(ctx, serviceProfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSpec", reflect.TypeOf((*MockServiceProfileClient)(nil).UpsertSpec), ctx, serviceProfile)
}

// DeleteService mocks base method.
func (m *MockServiceProfileClient) Delete(ctx context.Context, key client.ObjectKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockServiceProfileClientMockRecorder) Delete(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockServiceProfileClient)(nil).Delete), ctx, key)
}
