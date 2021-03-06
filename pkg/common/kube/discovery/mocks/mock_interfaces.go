// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_kubernetes_discovery is a generated GoMock package.
package mock_kubernetes_discovery

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	version "k8s.io/apimachinery/pkg/version"
)

// MockServerVersionClient is a mock of ServerVersionClient interface.
type MockServerVersionClient struct {
	ctrl     *gomock.Controller
	recorder *MockServerVersionClientMockRecorder
}

// MockServerVersionClientMockRecorder is the mock recorder for MockServerVersionClient.
type MockServerVersionClientMockRecorder struct {
	mock *MockServerVersionClient
}

// NewMockServerVersionClient creates a new mock instance.
func NewMockServerVersionClient(ctrl *gomock.Controller) *MockServerVersionClient {
	mock := &MockServerVersionClient{ctrl: ctrl}
	mock.recorder = &MockServerVersionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerVersionClient) EXPECT() *MockServerVersionClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockServerVersionClient) Get() (*version.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*version.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServerVersionClientMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServerVersionClient)(nil).Get))
}
