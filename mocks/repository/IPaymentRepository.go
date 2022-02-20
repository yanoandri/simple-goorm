// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/yanoandri/simple-goorm/models"
)

// IPaymentRepository is an autogenerated mock type for the IPaymentRepository type
type IPaymentRepository struct {
	mock.Mock
}

// CreatePayment provides a mock function with given fields: payment
func (_m *IPaymentRepository) CreatePayment(payment models.Payment) (int64, error) {
	ret := _m.Called(payment)

	var r0 int64
	if rf, ok := ret.Get(0).(func(models.Payment) int64); ok {
		r0 = rf(payment)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Payment) error); ok {
		r1 = rf(payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePayment provides a mock function with given fields: id
func (_m *IPaymentRepository) DeletePayment(id string) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectPaymentWIthId provides a mock function with given fields: id
func (_m *IPaymentRepository) SelectPaymentWIthId(id string) (models.Payment, error) {
	ret := _m.Called(id)

	var r0 models.Payment
	if rf, ok := ret.Get(0).(func(string) models.Payment); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Payment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePayment provides a mock function with given fields: id, payment
func (_m *IPaymentRepository) UpdatePayment(id string, payment models.Payment) (models.Payment, error) {
	ret := _m.Called(id, payment)

	var r0 models.Payment
	if rf, ok := ret.Get(0).(func(string, models.Payment) models.Payment); ok {
		r0 = rf(id, payment)
	} else {
		r0 = ret.Get(0).(models.Payment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.Payment) error); ok {
		r1 = rf(id, payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
