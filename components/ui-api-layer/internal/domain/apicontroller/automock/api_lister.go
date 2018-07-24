// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import v1alpha2 "github.com/kyma-project/kyma/components/api-controller/pkg/apis/gateway.kyma.cx/v1alpha2"

// apiLister is an autogenerated mock type for the apiLister type
type apiLister struct {
	mock.Mock
}

// List provides a mock function with given fields: environment, serviceName, hostname
func (_m *apiLister) List(environment string, serviceName *string, hostname *string) ([]*v1alpha2.Api, error) {
	ret := _m.Called(environment, serviceName, hostname)

	var r0 []*v1alpha2.Api
	if rf, ok := ret.Get(0).(func(string, *string, *string) []*v1alpha2.Api); ok {
		r0 = rf(environment, serviceName, hostname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha2.Api)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *string, *string) error); ok {
		r1 = rf(environment, serviceName, hostname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
