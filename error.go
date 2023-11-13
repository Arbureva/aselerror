// Package aselerror 这是一个自定义的 error 包
package aselerror

import (
	"strconv"
)

type AselError struct {
	cause string
	code  uint32
	err   error

	// level Degree of disaster
	level int8

	// output Are details exported when calling Error()? Some information can't be returned directly to the front-end via the interface
	output bool
}

func New(cause string) *AselError {
	return &AselError{cause: cause}
}

func NewCauseAndCode(cause string, code uint32) *AselError {
	return &AselError{cause: cause, code: code}
}

func (e *AselError) Wrap(err error) *AselError {
	e.err = err
	return e
}

func (e *AselError) GetCode() uint32 {
	return e.code
}

func (e *AselError) UpdateCause(cause string) *AselError {
	e.cause = cause
	return e
}

// Update Escalation of disaster level
func (e *AselError) Update(level int8) *AselError {
	e.level = level
	return e
}

func (e *AselError) Unwarp() error {
	return e.err
}

// HideDetails Automatically calling the Error() method outputs a code status code instead of a string or error, which hides the internal error message description.
func (e *AselError) HideDetails() *AselError {
	e.output = true
	return e
}

// ShowDetails deactivate sensitive status
func (e *AselError) ShowDetails() *AselError {
	e.output = false
	return e
}

func (e *AselError) Error() string {
	switch {
	case e.output:
		return strconv.Itoa(int(e.code))
	case e.cause != "" && e.err != nil:
		return e.cause + ": " + e.err.Error()
	case e.cause != "":
		return e.cause
	case e.err != nil:
		return e.err.Error()
	default:
		return "Unknown error"
	}
}
