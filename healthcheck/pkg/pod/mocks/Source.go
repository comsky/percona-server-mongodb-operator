// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import pod "github.com/percona/percona-server-mongodb-operator/healthcheck/pkg/pod"
import mock "github.com/stretchr/testify/mock"

// Source is an autogenerated mock type for the Source type
type Source struct {
	mock.Mock
}

// GetTasks provides a mock function with given fields: podName
func (_m *Source) GetTasks(podName string) ([]pod.Task, error) {
	ret := _m.Called(podName)

	var r0 []pod.Task
	if rf, ok := ret.Get(0).(func(string) []pod.Task); ok {
		r0 = rf(podName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pod.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(podName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Source) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Pods provides a mock function with given fields:
func (_m *Source) Pods() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// URL provides a mock function with given fields:
func (_m *Source) URL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}