//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/UserService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserServiceIface is a mock of UserServiceIface interface.
type MockUserServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceIfaceMockRecorder
}

// MockUserServiceIfaceMockRecorder is the mock recorder for MockUserServiceIface.
type MockUserServiceIfaceMockRecorder struct {
	mock *MockUserServiceIface
}

// NewMockUserServiceIface creates a new mock instance.
func NewMockUserServiceIface(ctrl *gomock.Controller) *MockUserServiceIface {
	mock := &MockUserServiceIface{ctrl: ctrl}
	mock.recorder = &MockUserServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceIface) EXPECT() *MockUserServiceIfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserServiceIface) CreateUser(p *CreateUserParams) (*CreateUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", p)
	ret0, _ := ret[0].(*CreateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceIfaceMockRecorder) CreateUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceIface)(nil).CreateUser), p)
}

// DeleteUser mocks base method.
func (m *MockUserServiceIface) DeleteUser(p *DeleteUserParams) (*DeleteUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", p)
	ret0, _ := ret[0].(*DeleteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceIfaceMockRecorder) DeleteUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServiceIface)(nil).DeleteUser), p)
}

// DisableUser mocks base method.
func (m *MockUserServiceIface) DisableUser(p *DisableUserParams) (*DisableUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUser", p)
	ret0, _ := ret[0].(*DisableUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableUser indicates an expected call of DisableUser.
func (mr *MockUserServiceIfaceMockRecorder) DisableUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUser", reflect.TypeOf((*MockUserServiceIface)(nil).DisableUser), p)
}

// EnableUser mocks base method.
func (m *MockUserServiceIface) EnableUser(p *EnableUserParams) (*EnableUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableUser", p)
	ret0, _ := ret[0].(*EnableUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableUser indicates an expected call of EnableUser.
func (mr *MockUserServiceIfaceMockRecorder) EnableUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUser", reflect.TypeOf((*MockUserServiceIface)(nil).EnableUser), p)
}

// GetUser mocks base method.
func (m *MockUserServiceIface) GetUser(p *GetUserParams) (*GetUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", p)
	ret0, _ := ret[0].(*GetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceIfaceMockRecorder) GetUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserServiceIface)(nil).GetUser), p)
}

// GetUserByID mocks base method.
func (m *MockUserServiceIface) GetUserByID(id string, opts ...OptionFunc) (*User, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByID", varargs...)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceIfaceMockRecorder) GetUserByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserServiceIface)(nil).GetUserByID), varargs...)
}

// GetUserKeys mocks base method.
func (m *MockUserServiceIface) GetUserKeys(p *GetUserKeysParams) (*GetUserKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserKeys", p)
	ret0, _ := ret[0].(*GetUserKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserKeys indicates an expected call of GetUserKeys.
func (mr *MockUserServiceIfaceMockRecorder) GetUserKeys(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserKeys", reflect.TypeOf((*MockUserServiceIface)(nil).GetUserKeys), p)
}

// GetVirtualMachineUserData mocks base method.
func (m *MockUserServiceIface) GetVirtualMachineUserData(p *GetVirtualMachineUserDataParams) (*GetVirtualMachineUserDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineUserData", p)
	ret0, _ := ret[0].(*GetVirtualMachineUserDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineUserData indicates an expected call of GetVirtualMachineUserData.
func (mr *MockUserServiceIfaceMockRecorder) GetVirtualMachineUserData(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineUserData", reflect.TypeOf((*MockUserServiceIface)(nil).GetVirtualMachineUserData), p)
}

// ListUsers mocks base method.
func (m *MockUserServiceIface) ListUsers(p *ListUsersParams) (*ListUsersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", p)
	ret0, _ := ret[0].(*ListUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockUserServiceIfaceMockRecorder) ListUsers(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserServiceIface)(nil).ListUsers), p)
}

// LockUser mocks base method.
func (m *MockUserServiceIface) LockUser(p *LockUserParams) (*LockUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockUser", p)
	ret0, _ := ret[0].(*LockUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockUser indicates an expected call of LockUser.
func (mr *MockUserServiceIfaceMockRecorder) LockUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockUser", reflect.TypeOf((*MockUserServiceIface)(nil).LockUser), p)
}

// NewCreateUserParams mocks base method.
func (m *MockUserServiceIface) NewCreateUserParams(account, email, firstname, lastname, password, username string) *CreateUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateUserParams", account, email, firstname, lastname, password, username)
	ret0, _ := ret[0].(*CreateUserParams)
	return ret0
}

// NewCreateUserParams indicates an expected call of NewCreateUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewCreateUserParams(account, email, firstname, lastname, password, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewCreateUserParams), account, email, firstname, lastname, password, username)
}

// NewDeleteUserParams mocks base method.
func (m *MockUserServiceIface) NewDeleteUserParams(id string) *DeleteUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteUserParams", id)
	ret0, _ := ret[0].(*DeleteUserParams)
	return ret0
}

// NewDeleteUserParams indicates an expected call of NewDeleteUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewDeleteUserParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewDeleteUserParams), id)
}

// NewDisableUserParams mocks base method.
func (m *MockUserServiceIface) NewDisableUserParams(id string) *DisableUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDisableUserParams", id)
	ret0, _ := ret[0].(*DisableUserParams)
	return ret0
}

// NewDisableUserParams indicates an expected call of NewDisableUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewDisableUserParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDisableUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewDisableUserParams), id)
}

// NewEnableUserParams mocks base method.
func (m *MockUserServiceIface) NewEnableUserParams(id string) *EnableUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableUserParams", id)
	ret0, _ := ret[0].(*EnableUserParams)
	return ret0
}

// NewEnableUserParams indicates an expected call of NewEnableUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewEnableUserParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewEnableUserParams), id)
}

// NewGetUserKeysParams mocks base method.
func (m *MockUserServiceIface) NewGetUserKeysParams(id string) *GetUserKeysParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGetUserKeysParams", id)
	ret0, _ := ret[0].(*GetUserKeysParams)
	return ret0
}

// NewGetUserKeysParams indicates an expected call of NewGetUserKeysParams.
func (mr *MockUserServiceIfaceMockRecorder) NewGetUserKeysParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetUserKeysParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewGetUserKeysParams), id)
}

// NewGetUserParams mocks base method.
func (m *MockUserServiceIface) NewGetUserParams(userapikey string) *GetUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGetUserParams", userapikey)
	ret0, _ := ret[0].(*GetUserParams)
	return ret0
}

// NewGetUserParams indicates an expected call of NewGetUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewGetUserParams(userapikey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewGetUserParams), userapikey)
}

// NewGetVirtualMachineUserDataParams mocks base method.
func (m *MockUserServiceIface) NewGetVirtualMachineUserDataParams(virtualmachineid string) *GetVirtualMachineUserDataParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGetVirtualMachineUserDataParams", virtualmachineid)
	ret0, _ := ret[0].(*GetVirtualMachineUserDataParams)
	return ret0
}

// NewGetVirtualMachineUserDataParams indicates an expected call of NewGetVirtualMachineUserDataParams.
func (mr *MockUserServiceIfaceMockRecorder) NewGetVirtualMachineUserDataParams(virtualmachineid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetVirtualMachineUserDataParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewGetVirtualMachineUserDataParams), virtualmachineid)
}

// NewListUsersParams mocks base method.
func (m *MockUserServiceIface) NewListUsersParams() *ListUsersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListUsersParams")
	ret0, _ := ret[0].(*ListUsersParams)
	return ret0
}

// NewListUsersParams indicates an expected call of NewListUsersParams.
func (mr *MockUserServiceIfaceMockRecorder) NewListUsersParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListUsersParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewListUsersParams))
}

// NewLockUserParams mocks base method.
func (m *MockUserServiceIface) NewLockUserParams(id string) *LockUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLockUserParams", id)
	ret0, _ := ret[0].(*LockUserParams)
	return ret0
}

// NewLockUserParams indicates an expected call of NewLockUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewLockUserParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLockUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewLockUserParams), id)
}

// NewRegisterUserKeysParams mocks base method.
func (m *MockUserServiceIface) NewRegisterUserKeysParams(id string) *RegisterUserKeysParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRegisterUserKeysParams", id)
	ret0, _ := ret[0].(*RegisterUserKeysParams)
	return ret0
}

// NewRegisterUserKeysParams indicates an expected call of NewRegisterUserKeysParams.
func (mr *MockUserServiceIfaceMockRecorder) NewRegisterUserKeysParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRegisterUserKeysParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewRegisterUserKeysParams), id)
}

// NewUpdateUserParams mocks base method.
func (m *MockUserServiceIface) NewUpdateUserParams(id string) *UpdateUserParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateUserParams", id)
	ret0, _ := ret[0].(*UpdateUserParams)
	return ret0
}

// NewUpdateUserParams indicates an expected call of NewUpdateUserParams.
func (mr *MockUserServiceIfaceMockRecorder) NewUpdateUserParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateUserParams", reflect.TypeOf((*MockUserServiceIface)(nil).NewUpdateUserParams), id)
}

// RegisterUserKeys mocks base method.
func (m *MockUserServiceIface) RegisterUserKeys(p *RegisterUserKeysParams) (*RegisterUserKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUserKeys", p)
	ret0, _ := ret[0].(*RegisterUserKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUserKeys indicates an expected call of RegisterUserKeys.
func (mr *MockUserServiceIfaceMockRecorder) RegisterUserKeys(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUserKeys", reflect.TypeOf((*MockUserServiceIface)(nil).RegisterUserKeys), p)
}

// UpdateUser mocks base method.
func (m *MockUserServiceIface) UpdateUser(p *UpdateUserParams) (*UpdateUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", p)
	ret0, _ := ret[0].(*UpdateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceIfaceMockRecorder) UpdateUser(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserServiceIface)(nil).UpdateUser), p)
}