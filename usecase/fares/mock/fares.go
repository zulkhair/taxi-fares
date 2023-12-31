// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/fares/fares.go

// Package mock_fares is a generated GoMock package.
package mock_fares

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	taxidata "github.com/zulkhair/taxi-fares/domain/taxidata"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// CalculateFares mocks base method.
func (m *MockUsecase) CalculateFares(taxiData []taxidata.TaxiData) *taxidata.Fares {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateFares", taxiData)
	ret0, _ := ret[0].(*taxidata.Fares)
	return ret0
}

// CalculateFares indicates an expected call of CalculateFares.
func (mr *MockUsecaseMockRecorder) CalculateFares(taxiData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateFares", reflect.TypeOf((*MockUsecase)(nil).CalculateFares), taxiData)
}
