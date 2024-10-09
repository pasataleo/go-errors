package errors

// ErrorCode represents a code that can be used to identify the type of error.
type ErrorCode string

const (
	// ErrorCodeUnknown is the default error code returned when an error hasn't been assigned a code.
	ErrorCodeUnknown ErrorCode = "ErrorCodeUnknown"

	// ErrorCodeOk is the error code returned when an error is nil.
	ErrorCodeOk ErrorCode = "ErrorCodeOk"

	// ErrorCodeWrapped is the error code returned when an error is wrapped by another error.
	ErrorCodeWrapped ErrorCode = "ErrorCodeWrapped"

	// ErrorCodeMulti is the error code returned when an error contains multiple errors.
	ErrorCodeMulti ErrorCode = "ErrorCodeMulti"
)

// Codeable is an interface that can be implemented by errors that have an error code.
type Codeable interface {
	error
	GetErrorCode() ErrorCode
}

// GetErrorCode returns the error code of the given error.
func GetErrorCode(err error) ErrorCode {
	if err == nil {
		return ErrorCodeOk
	}

	if codeable, ok := err.(Codeable); ok {
		code := codeable.GetErrorCode()
		if code == ErrorCodeWrapped {
			return GetErrorCode(Unwrap(err))
		}
		return code
	}
	return ErrorCodeUnknown
}
