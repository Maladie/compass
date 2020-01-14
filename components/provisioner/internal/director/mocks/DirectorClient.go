// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gqlschema "github.com/kyma-incubator/compass/components/provisioner/pkg/gqlschema"
import graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
import mock "github.com/stretchr/testify/mock"

// DirectorClient is an autogenerated mock type for the DirectorClient type
type DirectorClient struct {
	mock.Mock
}

// CreateRuntime provides a mock function with given fields: config, tenant
func (_m *DirectorClient) CreateRuntime(config *gqlschema.RuntimeInput, tenant string) (string, error) {
	ret := _m.Called(config, tenant)

	var r0 string
	if rf, ok := ret.Get(0).(func(*gqlschema.RuntimeInput, string) string); ok {
		r0 = rf(config, tenant)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gqlschema.RuntimeInput, string) error); ok {
		r1 = rf(config, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRuntime provides a mock function with given fields: id, tenant
func (_m *DirectorClient) DeleteRuntime(id string, tenant string) error {
	ret := _m.Called(id, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConnectionToken provides a mock function with given fields: id, tenant
func (_m *DirectorClient) GetConnectionToken(id string, tenant string) (graphql.OneTimeToken, error) {
	ret := _m.Called(id, tenant)

	var r0 graphql.OneTimeToken
	if rf, ok := ret.Get(0).(func(string, string) graphql.OneTimeToken); ok {
		r0 = rf(id, tenant)
	} else {
		r0 = ret.Get(0).(graphql.OneTimeToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
