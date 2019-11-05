// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/discovery_event_loop.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
)

// MockDiscoverySyncer is a mock of DiscoverySyncer interface
type MockDiscoverySyncer struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoverySyncerMockRecorder
}

// MockDiscoverySyncerMockRecorder is the mock recorder for MockDiscoverySyncer
type MockDiscoverySyncerMockRecorder struct {
	mock *MockDiscoverySyncer
}

// NewMockDiscoverySyncer creates a new mock instance
func NewMockDiscoverySyncer(ctrl *gomock.Controller) *MockDiscoverySyncer {
	mock := &MockDiscoverySyncer{ctrl: ctrl}
	mock.recorder = &MockDiscoverySyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoverySyncer) EXPECT() *MockDiscoverySyncerMockRecorder {
	return m.recorder
}

// Sync mocks base method
func (m *MockDiscoverySyncer) Sync(arg0 context.Context, arg1 *v1.DiscoverySnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockDiscoverySyncerMockRecorder) Sync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockDiscoverySyncer)(nil).Sync), arg0, arg1)
}
