// Code generated by MockGen. DO NOT EDIT.
// Source: test/service (interfaces: CustomerService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	model "test/model"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// GetCustomer mocks base method.
func (m *MockCustomerService) GetCustomer(arg0 int) (*model.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", arg0)
	ret0, _ := ret[0].(*model.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockCustomerServiceMockRecorder) GetCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockCustomerService)(nil).GetCustomer), arg0)
}

// GetCustomers mocks base method.
func (m *MockCustomerService) GetCustomers() ([]model.CustomerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomers")
	ret0, _ := ret[0].([]model.CustomerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomers indicates an expected call of GetCustomers.
func (mr *MockCustomerServiceMockRecorder) GetCustomers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomers", reflect.TypeOf((*MockCustomerService)(nil).GetCustomers))
}
