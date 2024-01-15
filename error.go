package aselerror

import (
	"fmt"
)

type AselError interface {
	GetCode() int
	GetLevel() int8
	GetMessage() string
	Error() string
}

type aerr struct {
	cause error
	code  int
	msg   string

	// level Degree of disaster
	level int8
}

type AselOption func(*aerr)

func WithCode(code int) AselOption {
	return func(c *aerr) {
		c.code = code
	}
}

func WithCause(err error) AselOption {
	return func(c *aerr) {
		c.cause = err
	}
}

func WithLevel(level int8) AselOption {
	return func(c *aerr) {
		c.level = level
	}
}

func New(message string, opts ...AselOption) AselError {
	var err aerr

	if len(message) != 0 {
		err.msg = message
	} else {
		err.msg = "unknow error"
	}

	for _, opt := range opts {
		opt(&err)
	}

	return err
}

func Wrap(err error, message string) AselError {
	if err == nil {
		return nil
	}

	return aerr{
		cause: err,
		msg:   message,
	}
}

func (e aerr) GetCode() int {
	return e.code
}

func (e aerr) GetLevel() int8 {
	return e.level
}

func (e aerr) GetMessage() string {
	return e.msg
}

func (e aerr) Unwrap() error {
	return e.cause
}

func (e aerr) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("aselerror cause: %s, message: %s, level: %d", e.cause, e.msg, e.level)
	} else {
		return fmt.Sprintf("aselerror cause: <nil>, message: %s, level: %d", e.msg, e.level)
	}
}
