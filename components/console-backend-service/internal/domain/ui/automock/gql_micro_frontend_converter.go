// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/kyma-project/kyma/common/microfrontend-client/pkg/apis/ui/v1alpha1"

// gqlMicroFrontendConverter is an autogenerated mock type for the gqlMicroFrontendConverter type
type gqlMicroFrontendConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlMicroFrontendConverter) ToGQL(in *v1alpha1.MicroFrontend) (*gqlschema.MicroFrontend, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.MicroFrontend
	if rf, ok := ret.Get(0).(func(*v1alpha1.MicroFrontend) *gqlschema.MicroFrontend); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.MicroFrontend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.MicroFrontend) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlMicroFrontendConverter) ToGQLs(in []*v1alpha1.MicroFrontend) ([]gqlschema.MicroFrontend, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.MicroFrontend
	if rf, ok := ret.Get(0).(func([]*v1alpha1.MicroFrontend) []gqlschema.MicroFrontend); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.MicroFrontend)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1alpha1.MicroFrontend) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}