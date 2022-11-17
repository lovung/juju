// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: UpgradeSeriesBackend,UpgradeSeriesMachine,UpgradeSeriesUnit)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/juju/juju/apiserver/common"
	model "github.com/juju/juju/core/model"
	status "github.com/juju/juju/core/status"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
)

// MockUpgradeSeriesBackend is a mock of UpgradeSeriesBackend interface.
type MockUpgradeSeriesBackend struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeSeriesBackendMockRecorder
}

// MockUpgradeSeriesBackendMockRecorder is the mock recorder for MockUpgradeSeriesBackend.
type MockUpgradeSeriesBackendMockRecorder struct {
	mock *MockUpgradeSeriesBackend
}

// NewMockUpgradeSeriesBackend creates a new mock instance.
func NewMockUpgradeSeriesBackend(ctrl *gomock.Controller) *MockUpgradeSeriesBackend {
	mock := &MockUpgradeSeriesBackend{ctrl: ctrl}
	mock.recorder = &MockUpgradeSeriesBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeSeriesBackend) EXPECT() *MockUpgradeSeriesBackendMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockUpgradeSeriesBackend) Machine(arg0 string) (common.UpgradeSeriesMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(common.UpgradeSeriesMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockUpgradeSeriesBackendMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockUpgradeSeriesBackend)(nil).Machine), arg0)
}

// Unit mocks base method.
func (m *MockUpgradeSeriesBackend) Unit(arg0 string) (common.UpgradeSeriesUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(common.UpgradeSeriesUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockUpgradeSeriesBackendMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockUpgradeSeriesBackend)(nil).Unit), arg0)
}

// MockUpgradeSeriesMachine is a mock of UpgradeSeriesMachine interface.
type MockUpgradeSeriesMachine struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeSeriesMachineMockRecorder
}

// MockUpgradeSeriesMachineMockRecorder is the mock recorder for MockUpgradeSeriesMachine.
type MockUpgradeSeriesMachineMockRecorder struct {
	mock *MockUpgradeSeriesMachine
}

// NewMockUpgradeSeriesMachine creates a new mock instance.
func NewMockUpgradeSeriesMachine(ctrl *gomock.Controller) *MockUpgradeSeriesMachine {
	mock := &MockUpgradeSeriesMachine{ctrl: ctrl}
	mock.recorder = &MockUpgradeSeriesMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeSeriesMachine) EXPECT() *MockUpgradeSeriesMachineMockRecorder {
	return m.recorder
}

// Base mocks base method.
func (m *MockUpgradeSeriesMachine) Base() state.Base {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Base")
	ret0, _ := ret[0].(state.Base)
	return ret0
}

// Base indicates an expected call of Base.
func (mr *MockUpgradeSeriesMachineMockRecorder) Base() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Base", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).Base))
}

// RemoveUpgradeSeriesLock mocks base method.
func (m *MockUpgradeSeriesMachine) RemoveUpgradeSeriesLock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUpgradeSeriesLock")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUpgradeSeriesLock indicates an expected call of RemoveUpgradeSeriesLock.
func (mr *MockUpgradeSeriesMachineMockRecorder) RemoveUpgradeSeriesLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUpgradeSeriesLock", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).RemoveUpgradeSeriesLock))
}

// SetInstanceStatus mocks base method.
func (m *MockUpgradeSeriesMachine) SetInstanceStatus(arg0 status.StatusInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInstanceStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInstanceStatus indicates an expected call of SetInstanceStatus.
func (mr *MockUpgradeSeriesMachineMockRecorder) SetInstanceStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstanceStatus", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).SetInstanceStatus), arg0)
}

// SetUpgradeSeriesStatus mocks base method.
func (m *MockUpgradeSeriesMachine) SetUpgradeSeriesStatus(arg0 model.UpgradeSeriesStatus, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUpgradeSeriesStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUpgradeSeriesStatus indicates an expected call of SetUpgradeSeriesStatus.
func (mr *MockUpgradeSeriesMachineMockRecorder) SetUpgradeSeriesStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpgradeSeriesStatus", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).SetUpgradeSeriesStatus), arg0, arg1)
}

// StartUpgradeSeriesUnitCompletion mocks base method.
func (m *MockUpgradeSeriesMachine) StartUpgradeSeriesUnitCompletion(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartUpgradeSeriesUnitCompletion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartUpgradeSeriesUnitCompletion indicates an expected call of StartUpgradeSeriesUnitCompletion.
func (mr *MockUpgradeSeriesMachineMockRecorder) StartUpgradeSeriesUnitCompletion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUpgradeSeriesUnitCompletion", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).StartUpgradeSeriesUnitCompletion), arg0)
}

// Units mocks base method.
func (m *MockUpgradeSeriesMachine) Units() ([]common.UpgradeSeriesUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]common.UpgradeSeriesUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockUpgradeSeriesMachineMockRecorder) Units() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).Units))
}

// UpdateMachineSeries mocks base method.
func (m *MockUpgradeSeriesMachine) UpdateMachineSeries(arg0 state.Base) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMachineSeries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMachineSeries indicates an expected call of UpdateMachineSeries.
func (mr *MockUpgradeSeriesMachineMockRecorder) UpdateMachineSeries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMachineSeries", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).UpdateMachineSeries), arg0)
}

// UpgradeSeriesStatus mocks base method.
func (m *MockUpgradeSeriesMachine) UpgradeSeriesStatus() (model.UpgradeSeriesStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesStatus")
	ret0, _ := ret[0].(model.UpgradeSeriesStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeSeriesStatus indicates an expected call of UpgradeSeriesStatus.
func (mr *MockUpgradeSeriesMachineMockRecorder) UpgradeSeriesStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesStatus", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).UpgradeSeriesStatus))
}

// UpgradeSeriesTarget mocks base method.
func (m *MockUpgradeSeriesMachine) UpgradeSeriesTarget() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesTarget")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeSeriesTarget indicates an expected call of UpgradeSeriesTarget.
func (mr *MockUpgradeSeriesMachineMockRecorder) UpgradeSeriesTarget() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesTarget", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).UpgradeSeriesTarget))
}

// UpgradeSeriesUnitStatuses mocks base method.
func (m *MockUpgradeSeriesMachine) UpgradeSeriesUnitStatuses() (map[string]state.UpgradeSeriesUnitStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesUnitStatuses")
	ret0, _ := ret[0].(map[string]state.UpgradeSeriesUnitStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeSeriesUnitStatuses indicates an expected call of UpgradeSeriesUnitStatuses.
func (mr *MockUpgradeSeriesMachineMockRecorder) UpgradeSeriesUnitStatuses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesUnitStatuses", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).UpgradeSeriesUnitStatuses))
}

// WatchUpgradeSeriesNotifications mocks base method.
func (m *MockUpgradeSeriesMachine) WatchUpgradeSeriesNotifications() (state.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpgradeSeriesNotifications")
	ret0, _ := ret[0].(state.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpgradeSeriesNotifications indicates an expected call of WatchUpgradeSeriesNotifications.
func (mr *MockUpgradeSeriesMachineMockRecorder) WatchUpgradeSeriesNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpgradeSeriesNotifications", reflect.TypeOf((*MockUpgradeSeriesMachine)(nil).WatchUpgradeSeriesNotifications))
}

// MockUpgradeSeriesUnit is a mock of UpgradeSeriesUnit interface.
type MockUpgradeSeriesUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeSeriesUnitMockRecorder
}

// MockUpgradeSeriesUnitMockRecorder is the mock recorder for MockUpgradeSeriesUnit.
type MockUpgradeSeriesUnitMockRecorder struct {
	mock *MockUpgradeSeriesUnit
}

// NewMockUpgradeSeriesUnit creates a new mock instance.
func NewMockUpgradeSeriesUnit(ctrl *gomock.Controller) *MockUpgradeSeriesUnit {
	mock := &MockUpgradeSeriesUnit{ctrl: ctrl}
	mock.recorder = &MockUpgradeSeriesUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeSeriesUnit) EXPECT() *MockUpgradeSeriesUnitMockRecorder {
	return m.recorder
}

// AssignedMachineId mocks base method.
func (m *MockUpgradeSeriesUnit) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockUpgradeSeriesUnitMockRecorder) AssignedMachineId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockUpgradeSeriesUnit)(nil).AssignedMachineId))
}

// SetUpgradeSeriesStatus mocks base method.
func (m *MockUpgradeSeriesUnit) SetUpgradeSeriesStatus(arg0 model.UpgradeSeriesStatus, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUpgradeSeriesStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUpgradeSeriesStatus indicates an expected call of SetUpgradeSeriesStatus.
func (mr *MockUpgradeSeriesUnitMockRecorder) SetUpgradeSeriesStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpgradeSeriesStatus", reflect.TypeOf((*MockUpgradeSeriesUnit)(nil).SetUpgradeSeriesStatus), arg0, arg1)
}

// Tag mocks base method.
func (m *MockUpgradeSeriesUnit) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockUpgradeSeriesUnitMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockUpgradeSeriesUnit)(nil).Tag))
}

// UpgradeSeriesStatus mocks base method.
func (m *MockUpgradeSeriesUnit) UpgradeSeriesStatus() (model.UpgradeSeriesStatus, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesStatus")
	ret0, _ := ret[0].(model.UpgradeSeriesStatus)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpgradeSeriesStatus indicates an expected call of UpgradeSeriesStatus.
func (mr *MockUpgradeSeriesUnitMockRecorder) UpgradeSeriesStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesStatus", reflect.TypeOf((*MockUpgradeSeriesUnit)(nil).UpgradeSeriesStatus))
}
