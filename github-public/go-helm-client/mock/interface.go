// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mockhelmclient is a generated GoMock package.
package mockhelmclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	helmclient "github.com/mittwald/go-helm-client"
	action "helm.sh/helm/v3/pkg/action"
	release "helm.sh/helm/v3/pkg/release"
	repo "helm.sh/helm/v3/pkg/repo"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddOrUpdateChartRepo mocks base method.
func (m *MockClient) AddOrUpdateChartRepo(entry repo.Entry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateChartRepo", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateChartRepo indicates an expected call of AddOrUpdateChartRepo.
func (mr *MockClientMockRecorder) AddOrUpdateChartRepo(entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateChartRepo", reflect.TypeOf((*MockClient)(nil).AddOrUpdateChartRepo), entry)
}

// GetRelease mocks base method.
func (m *MockClient) GetRelease(name string) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelease", name)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelease indicates an expected call of GetRelease.
func (mr *MockClientMockRecorder) GetRelease(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelease", reflect.TypeOf((*MockClient)(nil).GetRelease), name)
}

// GetReleaseValues mocks base method.
func (m *MockClient) GetReleaseValues(name string, allValues bool) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseValues", name, allValues)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReleaseValues indicates an expected call of GetReleaseValues.
func (mr *MockClientMockRecorder) GetReleaseValues(name, allValues interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseValues", reflect.TypeOf((*MockClient)(nil).GetReleaseValues), name, allValues)
}

// InstallOrUpgradeChart mocks base method.
func (m *MockClient) InstallOrUpgradeChart(ctx context.Context, spec *helmclient.ChartSpec) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallOrUpgradeChart", ctx, spec)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstallOrUpgradeChart indicates an expected call of InstallOrUpgradeChart.
func (mr *MockClientMockRecorder) InstallOrUpgradeChart(ctx, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallOrUpgradeChart", reflect.TypeOf((*MockClient)(nil).InstallOrUpgradeChart), ctx, spec)
}

// LintChart mocks base method.
func (m *MockClient) LintChart(spec *helmclient.ChartSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LintChart", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// LintChart indicates an expected call of LintChart.
func (mr *MockClientMockRecorder) LintChart(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LintChart", reflect.TypeOf((*MockClient)(nil).LintChart), spec)
}

// ListDeployedReleases mocks base method.
func (m *MockClient) ListDeployedReleases() ([]*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployedReleases")
	ret0, _ := ret[0].([]*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployedReleases indicates an expected call of ListDeployedReleases.
func (mr *MockClientMockRecorder) ListDeployedReleases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployedReleases", reflect.TypeOf((*MockClient)(nil).ListDeployedReleases))
}

// ListReleaseHistory mocks base method.
func (m *MockClient) ListReleaseHistory(name string, max int) ([]*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleaseHistory", name, max)
	ret0, _ := ret[0].([]*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleaseHistory indicates an expected call of ListReleaseHistory.
func (mr *MockClientMockRecorder) ListReleaseHistory(name, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseHistory", reflect.TypeOf((*MockClient)(nil).ListReleaseHistory), name, max)
}

// RollbackRelease mocks base method.
func (m *MockClient) RollbackRelease(spec *helmclient.ChartSpec, version int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackRelease", spec, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackRelease indicates an expected call of RollbackRelease.
func (mr *MockClientMockRecorder) RollbackRelease(spec, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackRelease", reflect.TypeOf((*MockClient)(nil).RollbackRelease), spec, version)
}

// SetDebugLog mocks base method.
func (m *MockClient) SetDebugLog(debugLog action.DebugLog) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDebugLog", debugLog)
}

// SetDebugLog indicates an expected call of SetDebugLog.
func (mr *MockClientMockRecorder) SetDebugLog(debugLog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDebugLog", reflect.TypeOf((*MockClient)(nil).SetDebugLog), debugLog)
}

// TemplateChart mocks base method.
func (m *MockClient) TemplateChart(spec *helmclient.ChartSpec) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateChart", spec)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TemplateChart indicates an expected call of TemplateChart.
func (mr *MockClientMockRecorder) TemplateChart(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateChart", reflect.TypeOf((*MockClient)(nil).TemplateChart), spec)
}

// UninstallRelease mocks base method.
func (m *MockClient) UninstallRelease(spec *helmclient.ChartSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallRelease", spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallRelease indicates an expected call of UninstallRelease.
func (mr *MockClientMockRecorder) UninstallRelease(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallRelease", reflect.TypeOf((*MockClient)(nil).UninstallRelease), spec)
}

// UninstallReleaseByName mocks base method.
func (m *MockClient) UninstallReleaseByName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallReleaseByName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallReleaseByName indicates an expected call of UninstallReleaseByName.
func (mr *MockClientMockRecorder) UninstallReleaseByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallReleaseByName", reflect.TypeOf((*MockClient)(nil).UninstallReleaseByName), name)
}

// UpdateChartRepos mocks base method.
func (m *MockClient) UpdateChartRepos() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChartRepos")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateChartRepos indicates an expected call of UpdateChartRepos.
func (mr *MockClientMockRecorder) UpdateChartRepos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChartRepos", reflect.TypeOf((*MockClient)(nil).UpdateChartRepos))
}
