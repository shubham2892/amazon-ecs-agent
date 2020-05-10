// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/logger/audit (interfaces: AuditLogger,InfoLogger)

// Package mock_audit is a generated GoMock package.
package mock_audit

import (
	reflect "reflect"

	request "github.com/aws/amazon-ecs-agent/agent/logger/audit/request"
	gomock "github.com/golang/mock/gomock"
)

// MockAuditLogger is a mock of AuditLogger interface
type MockAuditLogger struct {
	ctrl     *gomock.Controller
	recorder *MockAuditLoggerMockRecorder
}

// MockAuditLoggerMockRecorder is the mock recorder for MockAuditLogger
type MockAuditLoggerMockRecorder struct {
	mock *MockAuditLogger
}

// NewMockAuditLogger creates a new mock instance
func NewMockAuditLogger(ctrl *gomock.Controller) *MockAuditLogger {
	mock := &MockAuditLogger{ctrl: ctrl}
	mock.recorder = &MockAuditLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuditLogger) EXPECT() *MockAuditLoggerMockRecorder {
	return m.recorder
}

// GetCluster mocks base method
func (m *MockAuditLogger) GetCluster() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockAuditLoggerMockRecorder) GetCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockAuditLogger)(nil).GetCluster))
}

// GetContainerInstanceArn mocks base method
func (m *MockAuditLogger) GetContainerInstanceArn() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerInstanceArn")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContainerInstanceArn indicates an expected call of GetContainerInstanceArn
func (mr *MockAuditLoggerMockRecorder) GetContainerInstanceArn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerInstanceArn", reflect.TypeOf((*MockAuditLogger)(nil).GetContainerInstanceArn))
}

// Log mocks base method
func (m *MockAuditLogger) Log(arg0 request.LogRequest, arg1 int, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", arg0, arg1, arg2)
}

// Log indicates an expected call of Log
func (mr *MockAuditLoggerMockRecorder) Log(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockAuditLogger)(nil).Log), arg0, arg1, arg2)
}

// MockInfoLogger is a mock of InfoLogger interface
type MockInfoLogger struct {
	ctrl     *gomock.Controller
	recorder *MockInfoLoggerMockRecorder
}

// MockInfoLoggerMockRecorder is the mock recorder for MockInfoLogger
type MockInfoLoggerMockRecorder struct {
	mock *MockInfoLogger
}

// NewMockInfoLogger creates a new mock instance
func NewMockInfoLogger(ctrl *gomock.Controller) *MockInfoLogger {
	mock := &MockInfoLogger{ctrl: ctrl}
	mock.recorder = &MockInfoLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfoLogger) EXPECT() *MockInfoLoggerMockRecorder {
	return m.recorder
}

// Info mocks base method
func (m *MockInfoLogger) Info(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockInfoLoggerMockRecorder) Info(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInfoLogger)(nil).Info), arg0...)
}
