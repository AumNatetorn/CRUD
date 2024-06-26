// Code generated by MockGen. DO NOT EDIT.
// Source: app/crud/handler.go
//
// Generated by this command:
//
//	mockgen -package=crud -source=app/crud/handler.go -destination=app/crud/handler_mocks_test.go
//

// Package crud is a generated GoMock package.
package crud

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCustomerStorage is a mock of CustomerStorage interface.
type MockCustomerStorage struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerStorageMockRecorder
}

// MockCustomerStorageMockRecorder is the mock recorder for MockCustomerStorage.
type MockCustomerStorageMockRecorder struct {
	mock *MockCustomerStorage
}

// NewMockCustomerStorage creates a new mock instance.
func NewMockCustomerStorage(ctrl *gomock.Controller) *MockCustomerStorage {
	mock := &MockCustomerStorage{ctrl: ctrl}
	mock.recorder = &MockCustomerStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerStorage) EXPECT() *MockCustomerStorageMockRecorder {
	return m.recorder
}

// DeleteOne mocks base method.
func (m *MockCustomerStorage) DeleteOne(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockCustomerStorageMockRecorder) DeleteOne(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockCustomerStorage)(nil).DeleteOne), ctx, id)
}

// FindOne mocks base method.
func (m *MockCustomerStorage) FindOne(ctx context.Context, id int) (*Customers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, id)
	ret0, _ := ret[0].(*Customers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCustomerStorageMockRecorder) FindOne(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCustomerStorage)(nil).FindOne), ctx, id)
}

// InsertOne mocks base method.
func (m *MockCustomerStorage) InsertOne(ctx context.Context, req Customers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockCustomerStorageMockRecorder) InsertOne(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockCustomerStorage)(nil).InsertOne), ctx, req)
}

// UpdateOne mocks base method.
func (m *MockCustomerStorage) UpdateOne(ctx context.Context, req Customers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockCustomerStorageMockRecorder) UpdateOne(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockCustomerStorage)(nil).UpdateOne), ctx, req)
}
