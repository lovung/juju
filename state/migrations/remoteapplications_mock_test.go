// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state/migrations (interfaces: MigrationRemoteApplication,AllRemoteApplicationSource,StatusSource,RemoteApplicationSource,RemoteApplicationModel)

// Package migrations is a generated GoMock package.
package migrations

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v4 "github.com/juju/description/v4"
	v40 "github.com/juju/names/v4"
)

// MockMigrationRemoteApplication is a mock of MigrationRemoteApplication interface.
type MockMigrationRemoteApplication struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationRemoteApplicationMockRecorder
}

// MockMigrationRemoteApplicationMockRecorder is the mock recorder for MockMigrationRemoteApplication.
type MockMigrationRemoteApplicationMockRecorder struct {
	mock *MockMigrationRemoteApplication
}

// NewMockMigrationRemoteApplication creates a new mock instance.
func NewMockMigrationRemoteApplication(ctrl *gomock.Controller) *MockMigrationRemoteApplication {
	mock := &MockMigrationRemoteApplication{ctrl: ctrl}
	mock.recorder = &MockMigrationRemoteApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrationRemoteApplication) EXPECT() *MockMigrationRemoteApplicationMockRecorder {
	return m.recorder
}

// Bindings mocks base method.
func (m *MockMigrationRemoteApplication) Bindings() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bindings")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Bindings indicates an expected call of Bindings.
func (mr *MockMigrationRemoteApplicationMockRecorder) Bindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bindings", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).Bindings))
}

// Endpoints mocks base method.
func (m *MockMigrationRemoteApplication) Endpoints() ([]MigrationRemoteEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoints")
	ret0, _ := ret[0].([]MigrationRemoteEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Endpoints indicates an expected call of Endpoints.
func (mr *MockMigrationRemoteApplicationMockRecorder) Endpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoints", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).Endpoints))
}

// GlobalKey mocks base method.
func (m *MockMigrationRemoteApplication) GlobalKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GlobalKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GlobalKey indicates an expected call of GlobalKey.
func (mr *MockMigrationRemoteApplicationMockRecorder) GlobalKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GlobalKey", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).GlobalKey))
}

// IsConsumerProxy mocks base method.
func (m *MockMigrationRemoteApplication) IsConsumerProxy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConsumerProxy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConsumerProxy indicates an expected call of IsConsumerProxy.
func (mr *MockMigrationRemoteApplicationMockRecorder) IsConsumerProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConsumerProxy", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).IsConsumerProxy))
}

// Macaroon mocks base method.
func (m *MockMigrationRemoteApplication) Macaroon() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Macaroon")
	ret0, _ := ret[0].(string)
	return ret0
}

// Macaroon indicates an expected call of Macaroon.
func (mr *MockMigrationRemoteApplicationMockRecorder) Macaroon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Macaroon", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).Macaroon))
}

// OfferUUID mocks base method.
func (m *MockMigrationRemoteApplication) OfferUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// OfferUUID indicates an expected call of OfferUUID.
func (mr *MockMigrationRemoteApplicationMockRecorder) OfferUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferUUID", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).OfferUUID))
}

// SourceModel mocks base method.
func (m *MockMigrationRemoteApplication) SourceModel() v40.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourceModel")
	ret0, _ := ret[0].(v40.ModelTag)
	return ret0
}

// SourceModel indicates an expected call of SourceModel.
func (mr *MockMigrationRemoteApplicationMockRecorder) SourceModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourceModel", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).SourceModel))
}

// Spaces mocks base method.
func (m *MockMigrationRemoteApplication) Spaces() []MigrationRemoteSpace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spaces")
	ret0, _ := ret[0].([]MigrationRemoteSpace)
	return ret0
}

// Spaces indicates an expected call of Spaces.
func (mr *MockMigrationRemoteApplicationMockRecorder) Spaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spaces", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).Spaces))
}

// Tag mocks base method.
func (m *MockMigrationRemoteApplication) Tag() v40.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(v40.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockMigrationRemoteApplicationMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).Tag))
}

// URL mocks base method.
func (m *MockMigrationRemoteApplication) URL() (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// URL indicates an expected call of URL.
func (mr *MockMigrationRemoteApplicationMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockMigrationRemoteApplication)(nil).URL))
}

// MockAllRemoteApplicationSource is a mock of AllRemoteApplicationSource interface.
type MockAllRemoteApplicationSource struct {
	ctrl     *gomock.Controller
	recorder *MockAllRemoteApplicationSourceMockRecorder
}

// MockAllRemoteApplicationSourceMockRecorder is the mock recorder for MockAllRemoteApplicationSource.
type MockAllRemoteApplicationSourceMockRecorder struct {
	mock *MockAllRemoteApplicationSource
}

// NewMockAllRemoteApplicationSource creates a new mock instance.
func NewMockAllRemoteApplicationSource(ctrl *gomock.Controller) *MockAllRemoteApplicationSource {
	mock := &MockAllRemoteApplicationSource{ctrl: ctrl}
	mock.recorder = &MockAllRemoteApplicationSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllRemoteApplicationSource) EXPECT() *MockAllRemoteApplicationSourceMockRecorder {
	return m.recorder
}

// AllRemoteApplications mocks base method.
func (m *MockAllRemoteApplicationSource) AllRemoteApplications() ([]MigrationRemoteApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRemoteApplications")
	ret0, _ := ret[0].([]MigrationRemoteApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRemoteApplications indicates an expected call of AllRemoteApplications.
func (mr *MockAllRemoteApplicationSourceMockRecorder) AllRemoteApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRemoteApplications", reflect.TypeOf((*MockAllRemoteApplicationSource)(nil).AllRemoteApplications))
}

// MockStatusSource is a mock of StatusSource interface.
type MockStatusSource struct {
	ctrl     *gomock.Controller
	recorder *MockStatusSourceMockRecorder
}

// MockStatusSourceMockRecorder is the mock recorder for MockStatusSource.
type MockStatusSourceMockRecorder struct {
	mock *MockStatusSource
}

// NewMockStatusSource creates a new mock instance.
func NewMockStatusSource(ctrl *gomock.Controller) *MockStatusSource {
	mock := &MockStatusSource{ctrl: ctrl}
	mock.recorder = &MockStatusSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusSource) EXPECT() *MockStatusSourceMockRecorder {
	return m.recorder
}

// StatusArgs mocks base method.
func (m *MockStatusSource) StatusArgs(arg0 string) (v4.StatusArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusArgs", arg0)
	ret0, _ := ret[0].(v4.StatusArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusArgs indicates an expected call of StatusArgs.
func (mr *MockStatusSourceMockRecorder) StatusArgs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusArgs", reflect.TypeOf((*MockStatusSource)(nil).StatusArgs), arg0)
}

// MockRemoteApplicationSource is a mock of RemoteApplicationSource interface.
type MockRemoteApplicationSource struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteApplicationSourceMockRecorder
}

// MockRemoteApplicationSourceMockRecorder is the mock recorder for MockRemoteApplicationSource.
type MockRemoteApplicationSourceMockRecorder struct {
	mock *MockRemoteApplicationSource
}

// NewMockRemoteApplicationSource creates a new mock instance.
func NewMockRemoteApplicationSource(ctrl *gomock.Controller) *MockRemoteApplicationSource {
	mock := &MockRemoteApplicationSource{ctrl: ctrl}
	mock.recorder = &MockRemoteApplicationSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteApplicationSource) EXPECT() *MockRemoteApplicationSourceMockRecorder {
	return m.recorder
}

// AllRemoteApplications mocks base method.
func (m *MockRemoteApplicationSource) AllRemoteApplications() ([]MigrationRemoteApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRemoteApplications")
	ret0, _ := ret[0].([]MigrationRemoteApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRemoteApplications indicates an expected call of AllRemoteApplications.
func (mr *MockRemoteApplicationSourceMockRecorder) AllRemoteApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRemoteApplications", reflect.TypeOf((*MockRemoteApplicationSource)(nil).AllRemoteApplications))
}

// StatusArgs mocks base method.
func (m *MockRemoteApplicationSource) StatusArgs(arg0 string) (v4.StatusArgs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusArgs", arg0)
	ret0, _ := ret[0].(v4.StatusArgs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusArgs indicates an expected call of StatusArgs.
func (mr *MockRemoteApplicationSourceMockRecorder) StatusArgs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusArgs", reflect.TypeOf((*MockRemoteApplicationSource)(nil).StatusArgs), arg0)
}

// MockRemoteApplicationModel is a mock of RemoteApplicationModel interface.
type MockRemoteApplicationModel struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteApplicationModelMockRecorder
}

// MockRemoteApplicationModelMockRecorder is the mock recorder for MockRemoteApplicationModel.
type MockRemoteApplicationModelMockRecorder struct {
	mock *MockRemoteApplicationModel
}

// NewMockRemoteApplicationModel creates a new mock instance.
func NewMockRemoteApplicationModel(ctrl *gomock.Controller) *MockRemoteApplicationModel {
	mock := &MockRemoteApplicationModel{ctrl: ctrl}
	mock.recorder = &MockRemoteApplicationModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteApplicationModel) EXPECT() *MockRemoteApplicationModelMockRecorder {
	return m.recorder
}

// AddRemoteApplication mocks base method.
func (m *MockRemoteApplicationModel) AddRemoteApplication(arg0 v4.RemoteApplicationArgs) v4.RemoteApplication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteApplication", arg0)
	ret0, _ := ret[0].(v4.RemoteApplication)
	return ret0
}

// AddRemoteApplication indicates an expected call of AddRemoteApplication.
func (mr *MockRemoteApplicationModelMockRecorder) AddRemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteApplication", reflect.TypeOf((*MockRemoteApplicationModel)(nil).AddRemoteApplication), arg0)
}
