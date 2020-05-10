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
// Source: github.com/aws/amazon-ecs-agent/agent/ecscni (interfaces: CNIClient)

// Package mock_ecscni is a generated GoMock package.
package mock_ecscni

import (
	context "context"
	reflect "reflect"
	time "time"

	ecscni "github.com/aws/amazon-ecs-agent/agent/ecscni"
	current "github.com/containernetworking/cni/pkg/types/current"
	gomock "github.com/golang/mock/gomock"
)

// MockCNIClient is a mock of CNIClient interface
type MockCNIClient struct {
	ctrl     *gomock.Controller
	recorder *MockCNIClientMockRecorder
}

// MockCNIClientMockRecorder is the mock recorder for MockCNIClient
type MockCNIClientMockRecorder struct {
	mock *MockCNIClient
}

// NewMockCNIClient creates a new mock instance
func NewMockCNIClient(ctrl *gomock.Controller) *MockCNIClient {
	mock := &MockCNIClient{ctrl: ctrl}
	mock.recorder = &MockCNIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCNIClient) EXPECT() *MockCNIClientMockRecorder {
	return m.recorder
}

// Capabilities mocks base method
func (m *MockCNIClient) Capabilities(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Capabilities indicates an expected call of Capabilities
func (mr *MockCNIClientMockRecorder) Capabilities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockCNIClient)(nil).Capabilities), arg0)
}

// CleanupNS mocks base method
func (m *MockCNIClient) CleanupNS(arg0 context.Context, arg1 *ecscni.Config, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupNS", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupNS indicates an expected call of CleanupNS
func (mr *MockCNIClientMockRecorder) CleanupNS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupNS", reflect.TypeOf((*MockCNIClient)(nil).CleanupNS), arg0, arg1, arg2)
}

// ReleaseIPResource mocks base method
func (m *MockCNIClient) ReleaseIPResource(arg0 context.Context, arg1 *ecscni.Config, arg2 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseIPResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseIPResource indicates an expected call of ReleaseIPResource
func (mr *MockCNIClientMockRecorder) ReleaseIPResource(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseIPResource", reflect.TypeOf((*MockCNIClient)(nil).ReleaseIPResource), arg0, arg1, arg2)
}

// SetupNS mocks base method
func (m *MockCNIClient) SetupNS(arg0 context.Context, arg1 *ecscni.Config, arg2 time.Duration) (*current.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupNS", arg0, arg1, arg2)
	ret0, _ := ret[0].(*current.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupNS indicates an expected call of SetupNS
func (mr *MockCNIClientMockRecorder) SetupNS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNS", reflect.TypeOf((*MockCNIClient)(nil).SetupNS), arg0, arg1, arg2)
}

// Version mocks base method
func (m *MockCNIClient) Version(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockCNIClientMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCNIClient)(nil).Version), arg0)
}
