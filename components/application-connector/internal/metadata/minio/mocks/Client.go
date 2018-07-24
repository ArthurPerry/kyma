// Code generated by mockery v1.0.0
package mocks

import context "context"
import io "io"

import minio_go "github.com/minio/minio-go"
import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BucketExists provides a mock function with given fields: bucketName
func (_m *Client) BucketExists(bucketName string) (bool, error) {
	ret := _m.Called(bucketName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(bucketName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucketName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectWithContext provides a mock function with given fields: ctx, bucketName, objectName, opts
func (_m *Client) GetObjectWithContext(ctx context.Context, bucketName string, objectName string, opts minio_go.GetObjectOptions) (*minio_go.Object, error) {
	ret := _m.Called(ctx, bucketName, objectName, opts)

	var r0 *minio_go.Object
	if rf, ok := ret.Get(0).(func(context.Context, string, string, minio_go.GetObjectOptions) *minio_go.Object); ok {
		r0 = rf(ctx, bucketName, objectName, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*minio_go.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, minio_go.GetObjectOptions) error); ok {
		r1 = rf(ctx, bucketName, objectName, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeBucket provides a mock function with given fields: bucketName, location
func (_m *Client) MakeBucket(bucketName string, location string) error {
	ret := _m.Called(bucketName, location)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(bucketName, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutObjectWithContext provides a mock function with given fields: ctx, bucketName, objectName, reader, objectSize, opts
func (_m *Client) PutObjectWithContext(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio_go.PutObjectOptions) (int64, error) {
	ret := _m.Called(ctx, bucketName, objectName, reader, objectSize, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, int64, minio_go.PutObjectOptions) int64); ok {
		r0 = rf(ctx, bucketName, objectName, reader, objectSize, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader, int64, minio_go.PutObjectOptions) error); ok {
		r1 = rf(ctx, bucketName, objectName, reader, objectSize, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveObject provides a mock function with given fields: bucketName, objectName
func (_m *Client) RemoveObject(bucketName string, objectName string) error {
	ret := _m.Called(bucketName, objectName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(bucketName, objectName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
