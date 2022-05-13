// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/instill-ai/pipeline-backend/pkg/repository (interfaces: Repository)

// Package service_test is a generated GoMock package.
package service_test

import (
	reflect "reflect"

	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	datamodel "github.com/instill-ai/pipeline-backend/pkg/datamodel"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreatePipeline mocks base method.
func (m *MockRepository) CreatePipeline(arg0 *datamodel.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockRepositoryMockRecorder) CreatePipeline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockRepository)(nil).CreatePipeline), arg0)
}

// DeletePipeline mocks base method.
func (m *MockRepository) DeletePipeline(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockRepositoryMockRecorder) DeletePipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockRepository)(nil).DeletePipeline), arg0, arg1)
}

// GetPipelineByID mocks base method.
func (m *MockRepository) GetPipelineByID(arg0, arg1 string, arg2 bool) (*datamodel.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*datamodel.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineByID indicates an expected call of GetPipelineByID.
func (mr *MockRepositoryMockRecorder) GetPipelineByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineByID", reflect.TypeOf((*MockRepository)(nil).GetPipelineByID), arg0, arg1, arg2)
}

// GetPipelineByUID mocks base method.
func (m *MockRepository) GetPipelineByUID(arg0 uuid.UUID, arg1 string, arg2 bool) (*datamodel.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineByUID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*datamodel.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineByUID indicates an expected call of GetPipelineByUID.
func (mr *MockRepositoryMockRecorder) GetPipelineByUID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineByUID", reflect.TypeOf((*MockRepository)(nil).GetPipelineByUID), arg0, arg1, arg2)
}

// ListPipeline mocks base method.
func (m *MockRepository) ListPipeline(arg0 string, arg1 int, arg2 string, arg3 bool) ([]datamodel.Pipeline, int64, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipeline", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]datamodel.Pipeline)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListPipeline indicates an expected call of ListPipeline.
func (mr *MockRepositoryMockRecorder) ListPipeline(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipeline", reflect.TypeOf((*MockRepository)(nil).ListPipeline), arg0, arg1, arg2, arg3)
}

// UpdatePipeline mocks base method.
func (m *MockRepository) UpdatePipeline(arg0, arg1 string, arg2 *datamodel.Pipeline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipeline indicates an expected call of UpdatePipeline.
func (mr *MockRepositoryMockRecorder) UpdatePipeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockRepository)(nil).UpdatePipeline), arg0, arg1, arg2)
}

// UpdatePipelineID mocks base method.
func (m *MockRepository) UpdatePipelineID(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipelineID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipelineID indicates an expected call of UpdatePipelineID.
func (mr *MockRepositoryMockRecorder) UpdatePipelineID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipelineID", reflect.TypeOf((*MockRepository)(nil).UpdatePipelineID), arg0, arg1, arg2)
}

// UpdatePipelineState mocks base method.
func (m *MockRepository) UpdatePipelineState(arg0, arg1 string, arg2 datamodel.PipelineState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipelineState", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePipelineState indicates an expected call of UpdatePipelineState.
func (mr *MockRepositoryMockRecorder) UpdatePipelineState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipelineState", reflect.TypeOf((*MockRepository)(nil).UpdatePipelineState), arg0, arg1, arg2)
}