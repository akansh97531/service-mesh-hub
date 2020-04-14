// Code generated by MockGen. DO NOT EDIT.
// Source: ./unstructured_kube_client.go

// Package mock_kube is a generated GoMock package.
package mock_kube

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resource "k8s.io/cli-runtime/pkg/resource"
)

// MockUnstructuredKubeClient is a mock of UnstructuredKubeClient interface.
type MockUnstructuredKubeClient struct {
	ctrl     *gomock.Controller
	recorder *MockUnstructuredKubeClientMockRecorder
}

// MockUnstructuredKubeClientMockRecorder is the mock recorder for MockUnstructuredKubeClient.
type MockUnstructuredKubeClientMockRecorder struct {
	mock *MockUnstructuredKubeClient
}

// NewMockUnstructuredKubeClient creates a new mock instance.
func NewMockUnstructuredKubeClient(ctrl *gomock.Controller) *MockUnstructuredKubeClient {
	mock := &MockUnstructuredKubeClient{ctrl: ctrl}
	mock.recorder = &MockUnstructuredKubeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnstructuredKubeClient) EXPECT() *MockUnstructuredKubeClientMockRecorder {
	return m.recorder
}

// BuildResources mocks base method.
func (m *MockUnstructuredKubeClient) BuildResources(namespace, manifest string) ([]*resource.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildResources", namespace, manifest)
	ret0, _ := ret[0].([]*resource.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildResources indicates an expected call of BuildResources.
func (mr *MockUnstructuredKubeClientMockRecorder) BuildResources(namespace, manifest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildResources", reflect.TypeOf((*MockUnstructuredKubeClient)(nil).BuildResources), namespace, manifest)
}

// CreateMeshWorkload mocks base method.
func (m *MockUnstructuredKubeClient) Create(namespace string, resources []*resource.Info) ([]*resource.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeshWorkload", namespace, resources)
	ret0, _ := ret[0].([]*resource.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMeshWorkload indicates an expected call of CreateMeshWorkload.
func (mr *MockUnstructuredKubeClientMockRecorder) Create(namespace, resources interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeshWorkload", reflect.TypeOf((*MockUnstructuredKubeClient)(nil).Create), namespace, resources)
}

// DeleteService mocks base method.
func (m *MockUnstructuredKubeClient) Delete(namespace string, resources []*resource.Info) ([]*resource.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", namespace, resources)
	ret0, _ := ret[0].([]*resource.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockUnstructuredKubeClientMockRecorder) Delete(namespace, resources interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockUnstructuredKubeClient)(nil).Delete), namespace, resources)
}
