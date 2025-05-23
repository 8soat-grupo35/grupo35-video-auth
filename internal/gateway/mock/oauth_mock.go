// Code generated by MockGen. DO NOT EDIT.
// Source: oauth.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
)

// MockIOAuthConfig is a mock of IOAuthConfig interface.
type MockIOAuthConfig struct {
	ctrl     *gomock.Controller
	recorder *MockIOAuthConfigMockRecorder
}

// MockIOAuthConfigMockRecorder is the mock recorder for MockIOAuthConfig.
type MockIOAuthConfigMockRecorder struct {
	mock *MockIOAuthConfig
}

// NewMockIOAuthConfig creates a new mock instance.
func NewMockIOAuthConfig(ctrl *gomock.Controller) *MockIOAuthConfig {
	mock := &MockIOAuthConfig{ctrl: ctrl}
	mock.recorder = &MockIOAuthConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOAuthConfig) EXPECT() *MockIOAuthConfigMockRecorder {
	return m.recorder
}

// AuthCodeURL mocks base method.
func (m *MockIOAuthConfig) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{state}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthCodeURL", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthCodeURL indicates an expected call of AuthCodeURL.
func (mr *MockIOAuthConfigMockRecorder) AuthCodeURL(state interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{state}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCodeURL", reflect.TypeOf((*MockIOAuthConfig)(nil).AuthCodeURL), varargs...)
}

// Exchange mocks base method.
func (m *MockIOAuthConfig) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, code}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exchange", varargs...)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange.
func (mr *MockIOAuthConfigMockRecorder) Exchange(ctx, code interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, code}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockIOAuthConfig)(nil).Exchange), varargs...)
}

// Init mocks base method.
func (m *MockIOAuthConfig) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockIOAuthConfigMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockIOAuthConfig)(nil).Init))
}
