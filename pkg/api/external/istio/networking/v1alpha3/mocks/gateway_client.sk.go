// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/istio/networking/v1alpha3/gateway_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/solo-io/mesh-projects/pkg/api/external/istio/networking/v1alpha3"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockGatewayWatcher is a mock of GatewayWatcher interface
type MockGatewayWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayWatcherMockRecorder
}

// MockGatewayWatcherMockRecorder is the mock recorder for MockGatewayWatcher
type MockGatewayWatcherMockRecorder struct {
	mock *MockGatewayWatcher
}

// NewMockGatewayWatcher creates a new mock instance
func NewMockGatewayWatcher(ctrl *gomock.Controller) *MockGatewayWatcher {
	mock := &MockGatewayWatcher{ctrl: ctrl}
	mock.recorder = &MockGatewayWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewayWatcher) EXPECT() *MockGatewayWatcherMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockGatewayWatcher) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha3.GatewayList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha3.GatewayList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockGatewayWatcherMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockGatewayWatcher)(nil).Watch), namespace, opts)
}

// MockGatewayClient is a mock of GatewayClient interface
type MockGatewayClient struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClientMockRecorder
}

// MockGatewayClientMockRecorder is the mock recorder for MockGatewayClient
type MockGatewayClientMockRecorder struct {
	mock *MockGatewayClient
}

// NewMockGatewayClient creates a new mock instance
func NewMockGatewayClient(ctrl *gomock.Controller) *MockGatewayClient {
	mock := &MockGatewayClient{ctrl: ctrl}
	mock.recorder = &MockGatewayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGatewayClient) EXPECT() *MockGatewayClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockGatewayClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockGatewayClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockGatewayClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockGatewayClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockGatewayClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockGatewayClient)(nil).Register))
}

// Read mocks base method
func (m *MockGatewayClient) Read(namespace, name string, opts clients.ReadOpts) (*v1alpha3.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", namespace, name, opts)
	ret0, _ := ret[0].(*v1alpha3.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockGatewayClientMockRecorder) Read(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockGatewayClient)(nil).Read), namespace, name, opts)
}

// Write mocks base method
func (m *MockGatewayClient) Write(resource *v1alpha3.Gateway, opts clients.WriteOpts) (*v1alpha3.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*v1alpha3.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockGatewayClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockGatewayClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockGatewayClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGatewayClientMockRecorder) Delete(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGatewayClient)(nil).Delete), namespace, name, opts)
}

// List mocks base method
func (m *MockGatewayClient) List(namespace string, opts clients.ListOpts) (v1alpha3.GatewayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace, opts)
	ret0, _ := ret[0].(v1alpha3.GatewayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockGatewayClientMockRecorder) List(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGatewayClient)(nil).List), namespace, opts)
}

// Watch mocks base method
func (m *MockGatewayClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha3.GatewayList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha3.GatewayList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockGatewayClientMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockGatewayClient)(nil).Watch), namespace, opts)
}
