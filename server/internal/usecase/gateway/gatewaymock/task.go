// Code generated by MockGen. DO NOT EDIT.
// Source: ./task.go

// Package gatewaymock is a generated GoMock package.
package gatewaymock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	task "github.com/reearth/reearth-cms/server/pkg/task"
)

// MockTaskRunner is a mock of TaskRunner interface.
type MockTaskRunner struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRunnerMockRecorder
}

// MockTaskRunnerMockRecorder is the mock recorder for MockTaskRunner.
type MockTaskRunnerMockRecorder struct {
	mock *MockTaskRunner
}

// NewMockTaskRunner creates a new mock instance.
func NewMockTaskRunner(ctrl *gomock.Controller) *MockTaskRunner {
	mock := &MockTaskRunner{ctrl: ctrl}
	mock.recorder = &MockTaskRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRunner) EXPECT() *MockTaskRunnerMockRecorder {
	return m.recorder
}

// Retry mocks base method.
func (m *MockTaskRunner) Retry(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retry", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Retry indicates an expected call of Retry.
func (mr *MockTaskRunnerMockRecorder) Retry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retry", reflect.TypeOf((*MockTaskRunner)(nil).Retry), arg0, arg1)
}

// Run mocks base method.
func (m *MockTaskRunner) Run(arg0 context.Context, arg1 task.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockTaskRunnerMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTaskRunner)(nil).Run), arg0, arg1)
}
