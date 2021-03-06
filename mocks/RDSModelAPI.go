// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RDSModelAPI is an autogenerated mock type for the RDSModelAPI type
type RDSModelAPI struct {
	mock.Mock
}

// GetRDSClusterForTags provides a mock function with given fields: repository, branch
func (_m *RDSModelAPI) GetRDSClusterForTags(repository string, branch string) (*string, *string, error) {
	ret := _m.Called(repository, branch)

	var r0 *string
	if rf, ok := ret.Get(0).(func(string, string) *string); ok {
		r0 = rf(repository, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 *string
	if rf, ok := ret.Get(1).(func(string, string) *string); ok {
		r1 = rf(repository, branch)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(repository, branch)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StartRDSCluster provides a mock function with given fields: clusterARN, clusterStatus
func (_m *RDSModelAPI) StartRDSCluster(clusterARN *string, clusterStatus *string) (bool, error) {
	ret := _m.Called(clusterARN, clusterStatus)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*string, *string) bool); ok {
		r0 = rf(clusterARN, clusterStatus)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(clusterARN, clusterStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopRDSCluster provides a mock function with given fields: clusterARN, clusterStatus
func (_m *RDSModelAPI) StopRDSCluster(clusterARN *string, clusterStatus *string) (bool, error) {
	ret := _m.Called(clusterARN, clusterStatus)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*string, *string) bool); ok {
		r0 = rf(clusterARN, clusterStatus)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(clusterARN, clusterStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
