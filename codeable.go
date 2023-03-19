package errors

type ErrorCode string

const (
	ErrorCodeUnknown ErrorCode = "ErrorCodeUnknown"
	ErrorCodeOk      ErrorCode = "ErrorCodeOk"
	ErrorCodeWrapped ErrorCode = "ErrorCodeWrapped"
	ErrorCodeMulti   ErrorCode = "ErrorCodeMulti"
)

type Codeable interface {
	error
	GetErrorCode() ErrorCode
}

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
