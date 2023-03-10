// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	models "words-app/models"

	mock "github.com/stretchr/testify/mock"
)

// IDbHandler is an autogenerated mock type for the IDbHandler type
type IDbHandler struct {
	mock.Mock
}

// AddOne provides a mock function with given fields: word
func (_m *IDbHandler) AddOne(word string) (models.WordResponse, error) {
	ret := _m.Called(word)

	var r0 models.WordResponse
	if rf, ok := ret.Get(0).(func(string) models.WordResponse); ok {
		r0 = rf(word)
	} else {
		r0 = ret.Get(0).(models.WordResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(word)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *IDbHandler) GetAll() ([]models.WordResponse, error) {
	ret := _m.Called()

	var r0 []models.WordResponse
	if rf, ok := ret.Get(0).(func() []models.WordResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.WordResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIDbHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDbHandler creates a new instance of IDbHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDbHandler(t mockConstructorTestingTNewIDbHandler) *IDbHandler {
	mock := &IDbHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
