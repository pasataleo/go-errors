package errors

import (
	"errors"
	"fmt"
)

// New returns an error with a new error code, optionally wrapping an existing error.
//
// For example: errors.New(nil, ErrorCodeNotFound, "resource not found")
// For example: errors.New(err, ErrorCodeNotFound, "resource not found")
func New(err error, code ErrorCode, text string) error {
	return &WrappingError{
		err:   errors.New(text),
		wraps: err,
		code:  code,
	}
}

// Newf returns an error with a new error code, optionally wrapping an existing error.
//
// This matches New, but with a formatted message.
func Newf(err error, code ErrorCode, format string, args ...any) error {
	return &WrappingError{
		err:   fmt.Errorf(format, args...),
		wraps: err,
		code:  code,
	}
}

// Wrap returns an error wrapping an existing error.
//
// For example: errors.Wrap(err, "failed to read file")
func Wrap(err error, text string) error {
	return &WrappingError{
		err:   errors.New(text),
		wraps: err,
		code:  ErrorCodeWrapped,
	}
}

// Wrapf returns an error wrapping an existing error.
//
// This matches Wrap, but with a formatted message.
func Wrapf(err error, format string, args ...any) error {
	return &WrappingError{
		err:   fmt.Errorf(format, args),
		wraps: err,
		code:  ErrorCodeWrapped,
	}
}

// Embed returns a new error with the supplied data embedded in the provided error.
func Embed[Data any](err error, key string, value Data) error {
	if e, ok := err.(*DataContainingError); ok {
		// if this error already contains data, add to it.
		e.data[key] = value
		return e
	}

	if e, ok := err.(*WrappingError); ok {
		// If this is already a wrapping error, convert it to a data containing error.
		return &DataContainingError{
			WrappingError: e,
			data: map[string]interface{}{
				key: value,
			},
		}
	}

	return &DataContainingError{
		WrappingError: &WrappingError{
			err:   err,
			wraps: nil,
			code:  GetErrorCode(err),
		},
		data: map[string]interface{}{
			key: value,
		},
	}
}

// Is validates the error has the supplied error code.
func Is(err error, code ErrorCode) bool {
	return GetErrorCode(err) == code
}
