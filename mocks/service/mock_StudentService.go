// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vipul-08/cassandra-api/service (interfaces: StudentService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/vipul-08/cassandra-api/domain"
	exception "github.com/vipul-08/cassandra-api/exception"
)

// MockStudentService is a mock of StudentService interface.
type MockStudentService struct {
	ctrl     *gomock.Controller
	recorder *MockStudentServiceMockRecorder
}

// MockStudentServiceMockRecorder is the mock recorder for MockStudentService.
type MockStudentServiceMockRecorder struct {
	mock *MockStudentService
}

// NewMockStudentService creates a new mock instance.
func NewMockStudentService(ctrl *gomock.Controller) *MockStudentService {
	mock := &MockStudentService{ctrl: ctrl}
	mock.recorder = &MockStudentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentService) EXPECT() *MockStudentServiceMockRecorder {
	return m.recorder
}

// DeleteStudent mocks base method.
func (m *MockStudentService) DeleteStudent(arg0 gocql.UUID) *exception.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", arg0)
	ret0, _ := ret[0].(*exception.AppError)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockStudentServiceMockRecorder) DeleteStudent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockStudentService)(nil).DeleteStudent), arg0)
}

// GetAllStudents mocks base method.
func (m *MockStudentService) GetAllStudents() ([]domain.Student, *exception.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStudents")
	ret0, _ := ret[0].([]domain.Student)
	ret1, _ := ret[1].(*exception.AppError)
	return ret0, ret1
}

// GetAllStudents indicates an expected call of GetAllStudents.
func (mr *MockStudentServiceMockRecorder) GetAllStudents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStudents", reflect.TypeOf((*MockStudentService)(nil).GetAllStudents))
}

// GetStudentById mocks base method.
func (m *MockStudentService) GetStudentById(arg0 gocql.UUID) (*domain.Student, *exception.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentById", arg0)
	ret0, _ := ret[0].(*domain.Student)
	ret1, _ := ret[1].(*exception.AppError)
	return ret0, ret1
}

// GetStudentById indicates an expected call of GetStudentById.
func (mr *MockStudentServiceMockRecorder) GetStudentById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentById", reflect.TypeOf((*MockStudentService)(nil).GetStudentById), arg0)
}

// InsertStudent mocks base method.
func (m *MockStudentService) InsertStudent(arg0 *domain.Student) (*domain.Student, *exception.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertStudent", arg0)
	ret0, _ := ret[0].(*domain.Student)
	ret1, _ := ret[1].(*exception.AppError)
	return ret0, ret1
}

// InsertStudent indicates an expected call of InsertStudent.
func (mr *MockStudentServiceMockRecorder) InsertStudent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertStudent", reflect.TypeOf((*MockStudentService)(nil).InsertStudent), arg0)
}

// UpdateStudent mocks base method.
func (m *MockStudentService) UpdateStudent(arg0 *domain.Student, arg1 gocql.UUID) (*domain.Student, *exception.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", arg0, arg1)
	ret0, _ := ret[0].(*domain.Student)
	ret1, _ := ret[1].(*exception.AppError)
	return ret0, ret1
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockStudentServiceMockRecorder) UpdateStudent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockStudentService)(nil).UpdateStudent), arg0, arg1)
}