// Code generated by MockGen. DO NOT EDIT.
// Source: controller/redis.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRedis is a mock of Redis interface.
type MockRedis struct {
	ctrl     *gomock.Controller
	recorder *MockRedisMockRecorder
}

// MockRedisMockRecorder is the mock recorder for MockRedis.
type MockRedisMockRecorder struct {
	mock *MockRedis
}

// NewMockRedis creates a new mock instance.
func NewMockRedis(ctrl *gomock.Controller) *MockRedis {
	mock := &MockRedis{ctrl: ctrl}
	mock.recorder = &MockRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedis) EXPECT() *MockRedisMockRecorder {
	return m.recorder
}

// CacheKey mocks base method.
func (m *MockRedis) CacheKey(obj any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheKey", obj)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CacheKey indicates an expected call of CacheKey.
func (mr *MockRedisMockRecorder) CacheKey(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheKey", reflect.TypeOf((*MockRedis)(nil).CacheKey), obj)
}

// Close mocks base method.
func (m *MockRedis) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedis)(nil).Close))
}

// Get mocks base method.
func (m *MockRedis) Get(ctx context.Context, key string, want any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key, want)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisMockRecorder) Get(ctx, key, want interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedis)(nil).Get), ctx, key, want)
}

// Save mocks base method.
func (m *MockRedis) Save(ctx context.Context, key string, ret any) error {
	m.ctrl.T.Helper()
	ret_2 := m.ctrl.Call(m, "Save", ctx, key, ret)
	ret0, _ := ret_2[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockRedisMockRecorder) Save(ctx, key, ret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRedis)(nil).Save), ctx, key, ret)
}