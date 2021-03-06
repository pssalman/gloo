// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/plugins/hcm (interfaces: HcmPlugin)

// Package mock_hcm is a generated GoMock package.
package mock_hcm

import (
	reflect "reflect"

	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	hcm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	plugins "github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

// MockHcmPlugin is a mock of HcmPlugin interface
type MockHcmPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockHcmPluginMockRecorder
}

// MockHcmPluginMockRecorder is the mock recorder for MockHcmPlugin
type MockHcmPluginMockRecorder struct {
	mock *MockHcmPlugin
}

// NewMockHcmPlugin creates a new mock instance
func NewMockHcmPlugin(ctrl *gomock.Controller) *MockHcmPlugin {
	mock := &MockHcmPlugin{ctrl: ctrl}
	mock.recorder = &MockHcmPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHcmPlugin) EXPECT() *MockHcmPluginMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockHcmPlugin) Init(arg0 plugins.InitParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockHcmPluginMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockHcmPlugin)(nil).Init), arg0)
}

// ProcessHcmSettings mocks base method
func (m *MockHcmPlugin) ProcessHcmSettings(arg0 *v1.ApiSnapshot, arg1 *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager, arg2 *hcm.HttpConnectionManagerSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessHcmSettings", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessHcmSettings indicates an expected call of ProcessHcmSettings
func (mr *MockHcmPluginMockRecorder) ProcessHcmSettings(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessHcmSettings", reflect.TypeOf((*MockHcmPlugin)(nil).ProcessHcmSettings), arg0, arg1, arg2)
}
