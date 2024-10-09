package errors

import (
	"fmt"
	"strings"
)

var (
	_ error    = (*multi)(nil)
	_ Codeable = (*multi)(nil)
)

type multi struct {
	errs []error
}

// Error implements the error interface.
func (m *multi) Error() string {
	var errs []string
	for _, err := range m.errs {
		errs = append(errs, err.Error())
	}
	return fmt.Sprintf("multierror: [%s]", strings.Join(errs, ","))
}

// GetErrorCode implements the Codeable interface.
func (m *multi) GetErrorCode() ErrorCode {
	return ErrorCodeMulti
}

// Append appends the given errors to the current error.
func Append(current error, next ...error) error {
	if current == nil {
		if len(next) == 0 {
			return nil
		}

		if len(next) == 1 {
			return next[0]
		}

		return &multi{
			errs: next,
		}
	}

	if multi, ok := current.(*multi); ok {
		multi.errs = append(multi.errs, next...)
		return multi
	}

	var errs []error
	errs = append(errs, current)
	errs = append(errs, next...)

	return &multi{
		errs: errs,
	}
}

// Expand will expand the provider error into multiple errors if it is a multi error. Otherwise, it will return a slice
// containing the single error.
func Expand(err error) []error {
	if err == nil {
		return nil
	}

	if multi, ok := err.(*multi); ok {
		return multi.errs
	}

	return []error{err}
}
