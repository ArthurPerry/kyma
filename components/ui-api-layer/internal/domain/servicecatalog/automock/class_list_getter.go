// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import pager "github.com/kyma-project/kyma/components/ui-api-layer/internal/pager"

import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// classListGetter is an autogenerated mock type for the classListGetter type
type classListGetter struct {
	mock.Mock
}

// Find provides a mock function with given fields: name
func (_m *classListGetter) Find(name string) (*v1beta1.ClusterServiceClass, error) {
	ret := _m.Called(name)

	var r0 *v1beta1.ClusterServiceClass
	if rf, ok := ret.Get(0).(func(string) *v1beta1.ClusterServiceClass); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ClusterServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByExternalName provides a mock function with given fields: externalName
func (_m *classListGetter) FindByExternalName(externalName string) (*v1beta1.ClusterServiceClass, error) {
	ret := _m.Called(externalName)

	var r0 *v1beta1.ClusterServiceClass
	if rf, ok := ret.Get(0).(func(string) *v1beta1.ClusterServiceClass); ok {
		r0 = rf(externalName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ClusterServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(externalName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: pagingParams
func (_m *classListGetter) List(pagingParams pager.PagingParams) ([]*v1beta1.ClusterServiceClass, error) {
	ret := _m.Called(pagingParams)

	var r0 []*v1beta1.ClusterServiceClass
	if rf, ok := ret.Get(0).(func(pager.PagingParams) []*v1beta1.ClusterServiceClass); ok {
		r0 = rf(pagingParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.ClusterServiceClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pager.PagingParams) error); ok {
		r1 = rf(pagingParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
