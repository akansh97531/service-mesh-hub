// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_mesh_workload is a generated GoMock package.
package mock_mesh_workload

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMeshWorkloadFinder is a mock of MeshWorkloadFinder interface.
type MockMeshWorkloadFinder struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadFinderMockRecorder
}

// MockMeshWorkloadFinderMockRecorder is the mock recorder for MockMeshWorkloadFinder.
type MockMeshWorkloadFinderMockRecorder struct {
	mock *MockMeshWorkloadFinder
}

// NewMockMeshWorkloadFinder creates a new mock instance.
func NewMockMeshWorkloadFinder(ctrl *gomock.Controller) *MockMeshWorkloadFinder {
	mock := &MockMeshWorkloadFinder{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshWorkloadFinder) EXPECT() *MockMeshWorkloadFinderMockRecorder {
	return m.recorder
}

// StartDiscovery mocks base method.
func (m *MockMeshWorkloadFinder) StartDiscovery(podEventWatcher controller.PodController, predicates []predicate.Predicate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDiscovery", podEventWatcher, predicates)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartDiscovery indicates an expected call of StartDiscovery.
func (mr *MockMeshWorkloadFinderMockRecorder) StartDiscovery(podEventWatcher, predicates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDiscovery", reflect.TypeOf((*MockMeshWorkloadFinder)(nil).StartDiscovery), podEventWatcher, predicates)
}

// CreateMeshWorkload mocks base method.
func (m *MockMeshWorkloadFinder) Create(obj *v10.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMeshWorkload", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMeshWorkload indicates an expected call of CreateMeshWorkload.
func (mr *MockMeshWorkloadFinderMockRecorder) Create(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMeshWorkload", reflect.TypeOf((*MockMeshWorkloadFinder)(nil).Create), obj)
}

// UpdateService mocks base method.
func (m *MockMeshWorkloadFinder) Update(old, new *v10.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", old, new)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockMeshWorkloadFinderMockRecorder) Update(old, new interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockMeshWorkloadFinder)(nil).Update), old, new)
}

// DeleteService mocks base method.
func (m *MockMeshWorkloadFinder) Delete(obj *v10.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockMeshWorkloadFinderMockRecorder) Delete(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockMeshWorkloadFinder)(nil).Delete), obj)
}

// GenericService mocks base method.
func (m *MockMeshWorkloadFinder) Generic(obj *v10.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenericService", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenericService indicates an expected call of GenericService.
func (mr *MockMeshWorkloadFinderMockRecorder) Generic(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenericService", reflect.TypeOf((*MockMeshWorkloadFinder)(nil).Generic), obj)
}

// MockOwnerFetcher is a mock of OwnerFetcher interface.
type MockOwnerFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerFetcherMockRecorder
}

// MockOwnerFetcherMockRecorder is the mock recorder for MockOwnerFetcher.
type MockOwnerFetcherMockRecorder struct {
	mock *MockOwnerFetcher
}

// NewMockOwnerFetcher creates a new mock instance.
func NewMockOwnerFetcher(ctrl *gomock.Controller) *MockOwnerFetcher {
	mock := &MockOwnerFetcher{ctrl: ctrl}
	mock.recorder = &MockOwnerFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOwnerFetcher) EXPECT() *MockOwnerFetcherMockRecorder {
	return m.recorder
}

// GetDeployment mocks base method.
func (m *MockOwnerFetcher) GetDeployment(ctx context.Context, pod *v10.Pod) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, pod)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockOwnerFetcherMockRecorder) GetDeployment(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockOwnerFetcher)(nil).GetDeployment), ctx, pod)
}

// MockMeshWorkloadScanner is a mock of MeshWorkloadScanner interface.
type MockMeshWorkloadScanner struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadScannerMockRecorder
}

// MockMeshWorkloadScannerMockRecorder is the mock recorder for MockMeshWorkloadScanner.
type MockMeshWorkloadScannerMockRecorder struct {
	mock *MockMeshWorkloadScanner
}

// NewMockMeshWorkloadScanner creates a new mock instance.
func NewMockMeshWorkloadScanner(ctrl *gomock.Controller) *MockMeshWorkloadScanner {
	mock := &MockMeshWorkloadScanner{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshWorkloadScanner) EXPECT() *MockMeshWorkloadScannerMockRecorder {
	return m.recorder
}

// ScanPod mocks base method.
func (m *MockMeshWorkloadScanner) ScanPod(arg0 context.Context, arg1 *v10.Pod) (*types.ResourceRef, v11.ObjectMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanPod", arg0, arg1)
	ret0, _ := ret[0].(*types.ResourceRef)
	ret1, _ := ret[1].(v11.ObjectMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ScanPod indicates an expected call of ScanPod.
func (mr *MockMeshWorkloadScannerMockRecorder) ScanPod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanPod", reflect.TypeOf((*MockMeshWorkloadScanner)(nil).ScanPod), arg0, arg1)
}
