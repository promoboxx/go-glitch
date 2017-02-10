package error

import (
	"fmt"

	"github.com/lib/pq"
)

// UnknownCode is used as the code when a real code could not be found
const UnknownCode = "UNKNOWN"

// DataError is a class of error that provides access to an error code as well as the originating error
type DataError interface {
	// Error satisfies the error interface
	Error() string
	// Inner returns the originating error
	Inner() error
	// Code returns a specific error code, this is meant to be a machine friendly string
	Code() string
}

type dataError struct {
	inner error
	code  string
	msg   string
}

func (d *dataError) Error() string {
	return fmt.Sprintf("Code: [%s] Message: [%s] Inner error: [%s]", d.code, d.msg, d.inner.Error())
}

func (d *dataError) Inner() error {
	return d.inner
}

func (d *dataError) Code() string {
	return d.code
}

// FromPQ will create a DataError from a postgres error
func FromPQ(inner error, msg string) DataError {
	if pqErr, ok := inner.(pq.Error); ok {
		return &dataError{inner: inner, code: string(pqErr.Code), msg: msg}
	}
	return &dataError{inner: inner, code: UnknownCode, msg: msg}
}

// FromHTTPProblem will create a DataError from an HTTPProblem
func FromHTTPProblem(inner error, msg string) DataError {
	if httpProblem, ok := inner.(HTTPProblem); ok {
		return &dataError{inner: inner, code: httpProblem.Code, msg: msg}
	}
	return &dataError{inner: inner, code: UnknownCode, msg: msg}
}

// NewDataError will create a DataError from the information provided
func NewDataError(inner error, code string, msg string) DataError {
	return &dataError{inner: inner, code: code, msg: msg}
}
