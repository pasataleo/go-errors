package errors

import "fmt"

var (
	_ error         = (*WrappingError)(nil)
	_ Codeable      = (*WrappingError)(nil)
	_ DataContainer = (*DataContainingError)(nil)
)

type WrappingError struct {
	err   error
	wraps error

	code ErrorCode
}

func (e *WrappingError) Error() string {
	if e == nil {
		return ""
	}

	if e.wraps == nil {
		return e.err.Error()
	}
	return fmt.Sprintf("%s (%s)", e.err.Error(), e.wraps.Error())
}

func (e *WrappingError) GetErrorCode() ErrorCode {
	if e == nil {
		return ErrorCodeOk
	}
	return e.code
}

func (e *WrappingError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.wraps
}

type DataContainingError struct {
	*WrappingError

	data interface{}
}

func (d *DataContainingError) GetEmbeddedData() interface{} {
	return d.data
}
