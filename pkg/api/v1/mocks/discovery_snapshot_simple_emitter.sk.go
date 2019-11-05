// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/discovery_snapshot_simple_emitter.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
)

// MockDiscoverySimpleEmitter is a mock of DiscoverySimpleEmitter interface
type MockDiscoverySimpleEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoverySimpleEmitterMockRecorder
}

// MockDiscoverySimpleEmitterMockRecorder is the mock recorder for MockDiscoverySimpleEmitter
type MockDiscoverySimpleEmitterMockRecorder struct {
	mock *MockDiscoverySimpleEmitter
}

// NewMockDiscoverySimpleEmitter creates a new mock instance
func NewMockDiscoverySimpleEmitter(ctrl *gomock.Controller) *MockDiscoverySimpleEmitter {
	mock := &MockDiscoverySimpleEmitter{ctrl: ctrl}
	mock.recorder = &MockDiscoverySimpleEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoverySimpleEmitter) EXPECT() *MockDiscoverySimpleEmitterMockRecorder {
	return m.recorder
}

// Snapshots mocks base method
func (m *MockDiscoverySimpleEmitter) Snapshots(ctx context.Context) (<-chan *v1.DiscoverySnapshot, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshots", ctx)
	ret0, _ := ret[0].(<-chan *v1.DiscoverySnapshot)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Snapshots indicates an expected call of Snapshots
func (mr *MockDiscoverySimpleEmitterMockRecorder) Snapshots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshots", reflect.TypeOf((*MockDiscoverySimpleEmitter)(nil).Snapshots), ctx)
}
