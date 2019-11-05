// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/mesh_group_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockMeshGroupReconciler is a mock of MeshGroupReconciler interface
type MockMeshGroupReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMeshGroupReconcilerMockRecorder
}

// MockMeshGroupReconcilerMockRecorder is the mock recorder for MockMeshGroupReconciler
type MockMeshGroupReconcilerMockRecorder struct {
	mock *MockMeshGroupReconciler
}

// NewMockMeshGroupReconciler creates a new mock instance
func NewMockMeshGroupReconciler(ctrl *gomock.Controller) *MockMeshGroupReconciler {
	mock := &MockMeshGroupReconciler{ctrl: ctrl}
	mock.recorder = &MockMeshGroupReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshGroupReconciler) EXPECT() *MockMeshGroupReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockMeshGroupReconciler) Reconcile(namespace string, desiredResources v1.MeshGroupList, transition v1.TransitionMeshGroupFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockMeshGroupReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockMeshGroupReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}
