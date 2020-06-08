// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package v1_test is a generated GoMock package.
package v1_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	model "go.octolab.org/ecosystem/tablo/internal/model"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockStorage) Create(arg0 context.Context, arg1 []model.Board) ([]model.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].([]model.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStorageMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStorage)(nil).Create), arg0, arg1)
}

// CreateBoard mocks base method
func (m *MockStorage) CreateBoard(arg0 context.Context, arg1 model.Board) (model.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoard", arg0, arg1)
	ret0, _ := ret[0].(model.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBoard indicates an expected call of CreateBoard
func (mr *MockStorageMockRecorder) CreateBoard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoard", reflect.TypeOf((*MockStorage)(nil).CreateBoard), arg0, arg1)
}

// FetchBoard mocks base method
func (m *MockStorage) FetchBoard(arg0 context.Context, arg1 model.ID) (model.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBoard", arg0, arg1)
	ret0, _ := ret[0].(model.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBoard indicates an expected call of FetchBoard
func (mr *MockStorageMockRecorder) FetchBoard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBoard", reflect.TypeOf((*MockStorage)(nil).FetchBoard), arg0, arg1)
}

// FetchBoards mocks base method
func (m *MockStorage) FetchBoards(arg0 context.Context, arg1 map[string]interface{}) ([]model.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBoards", arg0, arg1)
	ret0, _ := ret[0].([]model.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBoards indicates an expected call of FetchBoards
func (mr *MockStorageMockRecorder) FetchBoards(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBoards", reflect.TypeOf((*MockStorage)(nil).FetchBoards), arg0, arg1)
}

// UpdateBoard mocks base method
func (m *MockStorage) UpdateBoard(arg0 context.Context, arg1 model.Board) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBoard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBoard indicates an expected call of UpdateBoard
func (mr *MockStorageMockRecorder) UpdateBoard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBoard", reflect.TypeOf((*MockStorage)(nil).UpdateBoard), arg0, arg1)
}

// DeleteBoard mocks base method
func (m *MockStorage) DeleteBoard(arg0 context.Context, arg1 model.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBoard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBoard indicates an expected call of DeleteBoard
func (mr *MockStorageMockRecorder) DeleteBoard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBoard", reflect.TypeOf((*MockStorage)(nil).DeleteBoard), arg0, arg1)
}

// CreateColumn mocks base method
func (m *MockStorage) CreateColumn(arg0 context.Context, arg1 model.Column) (model.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateColumn", arg0, arg1)
	ret0, _ := ret[0].(model.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateColumn indicates an expected call of CreateColumn
func (mr *MockStorageMockRecorder) CreateColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateColumn", reflect.TypeOf((*MockStorage)(nil).CreateColumn), arg0, arg1)
}

// FetchColumn mocks base method
func (m *MockStorage) FetchColumn(arg0 context.Context, arg1 model.ID) (model.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchColumn", arg0, arg1)
	ret0, _ := ret[0].(model.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchColumn indicates an expected call of FetchColumn
func (mr *MockStorageMockRecorder) FetchColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchColumn", reflect.TypeOf((*MockStorage)(nil).FetchColumn), arg0, arg1)
}

// UpdateColumn mocks base method
func (m *MockStorage) UpdateColumn(arg0 context.Context, arg1 model.Column) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn
func (mr *MockStorageMockRecorder) UpdateColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockStorage)(nil).UpdateColumn), arg0, arg1)
}

// DeleteColumn mocks base method
func (m *MockStorage) DeleteColumn(arg0 context.Context, arg1 model.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteColumn indicates an expected call of DeleteColumn
func (mr *MockStorageMockRecorder) DeleteColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteColumn", reflect.TypeOf((*MockStorage)(nil).DeleteColumn), arg0, arg1)
}

// CreateCard mocks base method
func (m *MockStorage) CreateCard(arg0 context.Context, arg1 model.Card) (model.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", arg0, arg1)
	ret0, _ := ret[0].(model.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCard indicates an expected call of CreateCard
func (mr *MockStorageMockRecorder) CreateCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockStorage)(nil).CreateCard), arg0, arg1)
}

// FetchCard mocks base method
func (m *MockStorage) FetchCard(arg0 context.Context, arg1 model.ID) (model.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCard", arg0, arg1)
	ret0, _ := ret[0].(model.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCard indicates an expected call of FetchCard
func (mr *MockStorageMockRecorder) FetchCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCard", reflect.TypeOf((*MockStorage)(nil).FetchCard), arg0, arg1)
}

// UpdateCard mocks base method
func (m *MockStorage) UpdateCard(arg0 context.Context, arg1 model.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCard indicates an expected call of UpdateCard
func (mr *MockStorageMockRecorder) UpdateCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockStorage)(nil).UpdateCard), arg0, arg1)
}

// DeleteCard mocks base method
func (m *MockStorage) DeleteCard(arg0 context.Context, arg1 model.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard
func (mr *MockStorageMockRecorder) DeleteCard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockStorage)(nil).DeleteCard), arg0, arg1)
}
