package errors

import "fmt"

var (
	_ error         = (*WrappingError)(nil)
	_ Codeable      = (*WrappingError)(nil)
	_ Unwrappable   = (*WrappingError)(nil)
	_ DataContainer = (*DataContainingError)(nil)
)

// WrappingError is an error that can wrap another error.
type WrappingError struct {
	err   error
	wraps error

	code ErrorCode
}

// Error implements the error interface.
func (e *WrappingError) Error() string {
	if e == nil {
		return ""
	}

	if e.wraps == nil {
		return e.err.Error()
	}
	return fmt.Sprintf("%s (%s)", e.err.Error(), e.wraps.Error())
}

// GetErrorCode implements the Codeable interface.
func (e *WrappingError) GetErrorCode() ErrorCode {
	if e == nil {
		return ErrorCodeOk
	}
	return e.code
}

// Unwrap implements the Unwrappable interface.
func (e *WrappingError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.wraps
}

// DataContainingError is an extension to the WrappingError struct that contains embedded data.
type DataContainingError struct {
	*WrappingError

	data map[string]interface{}
}

// EmbeddedData implements the DataContainer interface.
func (d *DataContainingError) EmbeddedData() map[string]interface{} {
	return d.data
}
