// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/database/database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "paru.net/gosimpleapp/internal/database"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// InsertFoo mocks base method.
func (m *MockDatabase) InsertFoo(name string, value int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFoo", name, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertFoo indicates an expected call of InsertFoo.
func (mr *MockDatabaseMockRecorder) InsertFoo(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFoo", reflect.TypeOf((*MockDatabase)(nil).InsertFoo), name, value)
}

// SelectFoo mocks base method.
func (m *MockDatabase) SelectFoo(id string) (database.Foo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFoo", id)
	ret0, _ := ret[0].(database.Foo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFoo indicates an expected call of SelectFoo.
func (mr *MockDatabaseMockRecorder) SelectFoo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFoo", reflect.TypeOf((*MockDatabase)(nil).SelectFoo), id)
}
