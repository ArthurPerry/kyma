// Code generated by mockery v1.0.0
package automock

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"

import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// gqlBrokerConverter is an autogenerated mock type for the gqlBrokerConverter type
type gqlBrokerConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlBrokerConverter) ToGQL(in *v1beta1.ClusterServiceBroker) (*gqlschema.ServiceBroker, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ServiceBroker
	if rf, ok := ret.Get(0).(func(*v1beta1.ClusterServiceBroker) *gqlschema.ServiceBroker); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ServiceBroker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.ClusterServiceBroker) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlBrokerConverter) ToGQLs(in []*v1beta1.ClusterServiceBroker) ([]gqlschema.ServiceBroker, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ServiceBroker
	if rf, ok := ret.Get(0).(func([]*v1beta1.ClusterServiceBroker) []gqlschema.ServiceBroker); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ServiceBroker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1beta1.ClusterServiceBroker) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
