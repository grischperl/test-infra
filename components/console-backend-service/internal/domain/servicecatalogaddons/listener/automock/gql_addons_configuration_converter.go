// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"

// gqlAddonsConfigurationConverter is an autogenerated mock type for the gqlAddonsConfigurationConverter type
type gqlAddonsConfigurationConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: item
func (_m *gqlAddonsConfigurationConverter) ToGQL(item *v1alpha1.AddonsConfiguration) *gqlschema.AddonsConfiguration {
	ret := _m.Called(item)

	var r0 *gqlschema.AddonsConfiguration
	if rf, ok := ret.Get(0).(func(*v1alpha1.AddonsConfiguration) *gqlschema.AddonsConfiguration); ok {
		r0 = rf(item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.AddonsConfiguration)
		}
	}

	return r0
}