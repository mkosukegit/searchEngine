package error

import (
	"fmt"
	"search/src/domain/error/code"
)

type MyError struct {
	Code code.Code
	err  error
}

func NewMyError(code code.Code, err any) *MyError {
	var e error
	switch err := err.(type) {
	case error:
		e = err
	default:
		e = fmt.Errorf("%v", err)
	}
	return &MyError{
		Code: code,
		err:  e,
	}
}

func (e *MyError) Error() string {
	return e.err.Error()
}

func (e *MyError) String() string {
	return e.Error()
}

func (e *MyError) Unwrap() error {
	return e.err
}
