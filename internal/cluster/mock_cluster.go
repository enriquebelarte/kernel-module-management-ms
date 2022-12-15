// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package cluster is a generated GoMock package.
package cluster

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	types "k8s.io/apimachinery/pkg/types"
	v1 "open-cluster-management.io/api/cluster/v1"
)

// MockClusterAPI is a mock of ClusterAPI interface.
type MockClusterAPI struct {
	ctrl     *gomock.Controller
	recorder *MockClusterAPIMockRecorder
}

// MockClusterAPIMockRecorder is the mock recorder for MockClusterAPI.
type MockClusterAPIMockRecorder struct {
	mock *MockClusterAPI
}

// NewMockClusterAPI creates a new mock instance.
func NewMockClusterAPI(ctrl *gomock.Controller) *MockClusterAPI {
	mock := &MockClusterAPI{ctrl: ctrl}
	mock.recorder = &MockClusterAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterAPI) EXPECT() *MockClusterAPIMockRecorder {
	return m.recorder
}

// BuildAndSign mocks base method.
func (m *MockClusterAPI) BuildAndSign(ctx context.Context, mcm v1beta1.ManagedClusterModule, cluster v1.ManagedCluster) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAndSign", ctx, mcm, cluster)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildAndSign indicates an expected call of BuildAndSign.
func (mr *MockClusterAPIMockRecorder) BuildAndSign(ctx, mcm, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndSign", reflect.TypeOf((*MockClusterAPI)(nil).BuildAndSign), ctx, mcm, cluster)
}

// GarbageCollectBuilds mocks base method.
func (m *MockClusterAPI) GarbageCollectBuilds(ctx context.Context, mcm v1beta1.ManagedClusterModule) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GarbageCollectBuilds", ctx, mcm)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GarbageCollectBuilds indicates an expected call of GarbageCollectBuilds.
func (mr *MockClusterAPIMockRecorder) GarbageCollectBuilds(ctx, mcm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GarbageCollectBuilds", reflect.TypeOf((*MockClusterAPI)(nil).GarbageCollectBuilds), ctx, mcm)
}

// RequestedManagedClusterModule mocks base method.
func (m *MockClusterAPI) RequestedManagedClusterModule(ctx context.Context, namespacedName types.NamespacedName) (*v1beta1.ManagedClusterModule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestedManagedClusterModule", ctx, namespacedName)
	ret0, _ := ret[0].(*v1beta1.ManagedClusterModule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestedManagedClusterModule indicates an expected call of RequestedManagedClusterModule.
func (mr *MockClusterAPIMockRecorder) RequestedManagedClusterModule(ctx, namespacedName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestedManagedClusterModule", reflect.TypeOf((*MockClusterAPI)(nil).RequestedManagedClusterModule), ctx, namespacedName)
}

// SelectedManagedClusters mocks base method.
func (m *MockClusterAPI) SelectedManagedClusters(ctx context.Context, mcm *v1beta1.ManagedClusterModule) (*v1.ManagedClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectedManagedClusters", ctx, mcm)
	ret0, _ := ret[0].(*v1.ManagedClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectedManagedClusters indicates an expected call of SelectedManagedClusters.
func (mr *MockClusterAPIMockRecorder) SelectedManagedClusters(ctx, mcm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectedManagedClusters", reflect.TypeOf((*MockClusterAPI)(nil).SelectedManagedClusters), ctx, mcm)
}
