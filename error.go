package aselerror

import (
	"fmt"
)

type AselError struct {
	cause string
	code  uint32
	err   error

	// level Degree of disaster
	level int8

	// private Are details exported when calling Error()? Some information can't be returned directly to the front-end via the interface
	private bool
}

type AselOption func(*AselError)

func WithCode(code uint32) AselOption {
	return func(c *AselError) {
		c.code = code
	}
}

func WithError(err error) AselOption {
	return func(c *AselError) {
		c.err = err
	}
}

func WithLevel(level int8) AselOption {
	return func(c *AselError) {
		c.level = level
	}
}

func WithOutput(level int8) AselOption {
	return func(c *AselError) {
		c.level = level
	}
}

func New(cause string, opts ...AselOption) *AselError {
	err := &AselError{cause: cause}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func (e *AselError) GetCode() uint32 {
	return e.code
}

// HideDetails Automatically calling the Error() method outputs a code status code instead of a string or error, which hides the internal error message description.
func (e *AselError) HideDetails() *AselError {
	e.private = true
	return e
}

// ShowDetails deactivate sensitive status
func (e *AselError) ShowDetails() *AselError {
	e.private = false
	return e
}

func (e *AselError) Error() string {
	switch {
	case e.private:
		return fmt.Sprintf("aselerror unknow! code: %d", e.code)
	case e.err != nil:
		return fmt.Sprintf("aselerror cause: %s, detail: %s", e.cause, e.err.Error())
	default:
		return fmt.Sprintf("aselerror cause: %s, detail: <nil>", e.cause)
	}
}
