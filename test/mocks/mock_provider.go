// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubecost/cost-model/cloud (interfaces: Provider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	cloud "github.com/kubecost/cost-model/cloud"
	io "io"
	v1 "k8s.io/api/core/v1"
	url "net/url"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AddServiceKey mocks base method
func (m *MockProvider) AddServiceKey(arg0 url.Values) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServiceKey", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddServiceKey indicates an expected call of AddServiceKey
func (mr *MockProviderMockRecorder) AddServiceKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServiceKey", reflect.TypeOf((*MockProvider)(nil).AddServiceKey), arg0)
}

// AllNodePricing mocks base method
func (m *MockProvider) AllNodePricing() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllNodePricing")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllNodePricing indicates an expected call of AllNodePricing
func (mr *MockProviderMockRecorder) AllNodePricing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllNodePricing", reflect.TypeOf((*MockProvider)(nil).AllNodePricing))
}

// ClusterInfo mocks base method
func (m *MockProvider) ClusterInfo() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterInfo")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterInfo indicates an expected call of ClusterInfo
func (mr *MockProviderMockRecorder) ClusterInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterInfo", reflect.TypeOf((*MockProvider)(nil).ClusterInfo))
}

// DownloadPricingData mocks base method
func (m *MockProvider) DownloadPricingData() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadPricingData")
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadPricingData indicates an expected call of DownloadPricingData
func (mr *MockProviderMockRecorder) DownloadPricingData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadPricingData", reflect.TypeOf((*MockProvider)(nil).DownloadPricingData))
}

// ExternalAllocations mocks base method
func (m *MockProvider) ExternalAllocations(arg0, arg1, arg2 string) ([]*cloud.OutOfClusterAllocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalAllocations", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*cloud.OutOfClusterAllocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExternalAllocations indicates an expected call of ExternalAllocations
func (mr *MockProviderMockRecorder) ExternalAllocations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalAllocations", reflect.TypeOf((*MockProvider)(nil).ExternalAllocations), arg0, arg1, arg2)
}

// GetConfig mocks base method
func (m *MockProvider) GetConfig() (*cloud.CustomPricing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*cloud.CustomPricing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockProviderMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockProvider)(nil).GetConfig))
}

// GetDisks mocks base method
func (m *MockProvider) GetDisks() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDisks")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDisks indicates an expected call of GetDisks
func (mr *MockProviderMockRecorder) GetDisks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisks", reflect.TypeOf((*MockProvider)(nil).GetDisks))
}

// GetKey mocks base method
func (m *MockProvider) GetKey(arg0 map[string]string) cloud.Key {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", arg0)
	ret0, _ := ret[0].(cloud.Key)
	return ret0
}

// GetKey indicates an expected call of GetKey
func (mr *MockProviderMockRecorder) GetKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockProvider)(nil).GetKey), arg0)
}

// GetLocalStorageQuery mocks base method
func (m *MockProvider) GetLocalStorageQuery() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalStorageQuery")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocalStorageQuery indicates an expected call of GetLocalStorageQuery
func (mr *MockProviderMockRecorder) GetLocalStorageQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalStorageQuery", reflect.TypeOf((*MockProvider)(nil).GetLocalStorageQuery))
}

// GetManagementPlatform mocks base method
func (m *MockProvider) GetManagementPlatform() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagementPlatform")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagementPlatform indicates an expected call of GetManagementPlatform
func (mr *MockProviderMockRecorder) GetManagementPlatform() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagementPlatform", reflect.TypeOf((*MockProvider)(nil).GetManagementPlatform))
}

// GetPVKey mocks base method
func (m *MockProvider) GetPVKey(arg0 *v1.PersistentVolume, arg1 map[string]string) cloud.PVKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPVKey", arg0, arg1)
	ret0, _ := ret[0].(cloud.PVKey)
	return ret0
}

// GetPVKey indicates an expected call of GetPVKey
func (mr *MockProviderMockRecorder) GetPVKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPVKey", reflect.TypeOf((*MockProvider)(nil).GetPVKey), arg0, arg1)
}

// NodePricing mocks base method
func (m *MockProvider) NodePricing(arg0 cloud.Key) (*cloud.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodePricing", arg0)
	ret0, _ := ret[0].(*cloud.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodePricing indicates an expected call of NodePricing
func (mr *MockProviderMockRecorder) NodePricing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodePricing", reflect.TypeOf((*MockProvider)(nil).NodePricing), arg0)
}

// PVPricing mocks base method
func (m *MockProvider) PVPricing(arg0 cloud.PVKey) (*cloud.PV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PVPricing", arg0)
	ret0, _ := ret[0].(*cloud.PV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PVPricing indicates an expected call of PVPricing
func (mr *MockProviderMockRecorder) PVPricing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PVPricing", reflect.TypeOf((*MockProvider)(nil).PVPricing), arg0)
}

// UpdateConfig mocks base method
func (m *MockProvider) UpdateConfig(arg0 io.Reader, arg1 string) (*cloud.CustomPricing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfig", arg0, arg1)
	ret0, _ := ret[0].(*cloud.CustomPricing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockProviderMockRecorder) UpdateConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockProvider)(nil).UpdateConfig), arg0, arg1)
}
