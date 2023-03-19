package errors

import (
	"errors"
	"fmt"
)

func New(err error, code ErrorCode, text string) error {
	return &WrappingError{
		err:   errors.New(text),
		wraps: err,
		code:  code,
	}
}

func Newf(err error, code ErrorCode, format string, args ...any) error {
	return &WrappingError{
		err:   fmt.Errorf(format, args...),
		wraps: err,
		code:  code,
	}
}

func Wrap(err error, text string) error {
	return &WrappingError{
		err:   errors.New(text),
		wraps: err,
		code:  ErrorCodeWrapped,
	}
}

func Wrapf(err error, format string, args ...any) error {
	return &WrappingError{
		err:   fmt.Errorf(format, args),
		wraps: err,
		code:  ErrorCodeWrapped,
	}
}

func Embed[Data any](err error, data Data) error {
	if e, ok := err.(*WrappingError); ok {
		return &DataContainingError[Data]{
			WrappingError: e,
			data:          data,
		}
	}

	return &DataContainingError[Data]{
		WrappingError: &WrappingError{
			err:   err,
			wraps: nil,
			code:  GetErrorCode(err),
		},
		data: data,
	}
}

func Is(err error, code ErrorCode) bool {
	return GetErrorCode(err) == code
}
