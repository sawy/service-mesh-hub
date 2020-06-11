// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_session is a generated GoMock package.
package mock_session

import (
	session "github.com/aws/aws-sdk-go/aws/session"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAwsSessionStore is a mock of AwsSessionStore interface.
type MockAwsSessionStore struct {
	ctrl     *gomock.Controller
	recorder *MockAwsSessionStoreMockRecorder
}

// MockAwsSessionStoreMockRecorder is the mock recorder for MockAwsSessionStore.
type MockAwsSessionStoreMockRecorder struct {
	mock *MockAwsSessionStore
}

// NewMockAwsSessionStore creates a new mock instance.
func NewMockAwsSessionStore(ctrl *gomock.Controller) *MockAwsSessionStore {
	mock := &MockAwsSessionStore{ctrl: ctrl}
	mock.recorder = &MockAwsSessionStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsSessionStore) EXPECT() *MockAwsSessionStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAwsSessionStore) Add(accountId, region string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", accountId, region)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockAwsSessionStoreMockRecorder) Add(accountId, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAwsSessionStore)(nil).Add), accountId, region)
}

// Get mocks base method.
func (m *MockAwsSessionStore) Get(accountId, region string) (*session.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", accountId, region)
	ret0, _ := ret[0].(*session.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAwsSessionStoreMockRecorder) Get(accountId, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAwsSessionStore)(nil).Get), accountId, region)
}