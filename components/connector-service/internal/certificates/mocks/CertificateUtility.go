// Code generated by mockery v1.0.0
package mocks

import apperrors "github.com/kyma-project/kyma/components/connector-service/internal/apperrors"
import certificates "github.com/kyma-project/kyma/components/connector-service/internal/certificates"
import mock "github.com/stretchr/testify/mock"
import rsa "crypto/rsa"
import x509 "crypto/x509"

// CertificateUtility is an autogenerated mock type for the CertificateUtility type
type CertificateUtility struct {
	mock.Mock
}

// CheckCSRValues provides a mock function with given fields: csr, subject
func (_m *CertificateUtility) CheckCSRValues(csr *x509.CertificateRequest, subject certificates.CSRSubject) apperrors.AppError {
	ret := _m.Called(csr, subject)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(*x509.CertificateRequest, certificates.CSRSubject) apperrors.AppError); ok {
		r0 = rf(csr, subject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// LoadCSR provides a mock function with given fields: encodedData
func (_m *CertificateUtility) LoadCSR(encodedData string) (*x509.CertificateRequest, apperrors.AppError) {
	ret := _m.Called(encodedData)

	var r0 *x509.CertificateRequest
	if rf, ok := ret.Get(0).(func(string) *x509.CertificateRequest); ok {
		r0 = rf(encodedData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*x509.CertificateRequest)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(encodedData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// LoadCert provides a mock function with given fields: encodedData
func (_m *CertificateUtility) LoadCert(encodedData []byte) (*x509.Certificate, apperrors.AppError) {
	ret := _m.Called(encodedData)

	var r0 *x509.Certificate
	if rf, ok := ret.Get(0).(func([]byte) *x509.Certificate); ok {
		r0 = rf(encodedData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*x509.Certificate)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func([]byte) apperrors.AppError); ok {
		r1 = rf(encodedData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// LoadKey provides a mock function with given fields: encodedData
func (_m *CertificateUtility) LoadKey(encodedData []byte) (*rsa.PrivateKey, apperrors.AppError) {
	ret := _m.Called(encodedData)

	var r0 *rsa.PrivateKey
	if rf, ok := ret.Get(0).(func([]byte) *rsa.PrivateKey); ok {
		r0 = rf(encodedData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rsa.PrivateKey)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func([]byte) apperrors.AppError); ok {
		r1 = rf(encodedData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// SignWithCA provides a mock function with given fields: caCrt, csr, key
func (_m *CertificateUtility) SignWithCA(caCrt *x509.Certificate, csr *x509.CertificateRequest, key *rsa.PrivateKey) (string, apperrors.AppError) {
	ret := _m.Called(caCrt, csr, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(*x509.Certificate, *x509.CertificateRequest, *rsa.PrivateKey) string); ok {
		r0 = rf(caCrt, csr, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(*x509.Certificate, *x509.CertificateRequest, *rsa.PrivateKey) apperrors.AppError); ok {
		r1 = rf(caCrt, csr, key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
