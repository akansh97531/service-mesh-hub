// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/istio/networking/v1alpha3/service_entry_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha3 "github.com/solo-io/mesh-projects/pkg/api/external/istio/networking/v1alpha3"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockServiceEntryReconciler is a mock of ServiceEntryReconciler interface
type MockServiceEntryReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockServiceEntryReconcilerMockRecorder
}

// MockServiceEntryReconcilerMockRecorder is the mock recorder for MockServiceEntryReconciler
type MockServiceEntryReconcilerMockRecorder struct {
	mock *MockServiceEntryReconciler
}

// NewMockServiceEntryReconciler creates a new mock instance
func NewMockServiceEntryReconciler(ctrl *gomock.Controller) *MockServiceEntryReconciler {
	mock := &MockServiceEntryReconciler{ctrl: ctrl}
	mock.recorder = &MockServiceEntryReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceEntryReconciler) EXPECT() *MockServiceEntryReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockServiceEntryReconciler) Reconcile(namespace string, desiredResources v1alpha3.ServiceEntryList, transition v1alpha3.TransitionServiceEntryFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockServiceEntryReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockServiceEntryReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}
