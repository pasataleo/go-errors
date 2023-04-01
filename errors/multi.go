package errors

import (
	"fmt"
	"strings"
)

var (
	_ error         = (*multi)(nil)
	_ Codeable      = (*multi)(nil)
	_ DataContainer = (*multi)(nil)
)

type multi struct {
	errs []error
}

func (m *multi) Error() string {
	var errs []string
	for _, err := range m.errs {
		errs = append(errs, err.Error())
	}
	return fmt.Sprintf("multierror: [%s]", strings.Join(errs, ","))
}

func (m *multi) GetErrorCode() ErrorCode {
	return ErrorCodeMulti
}

func (m *multi) GetEmbeddedData() interface{} {
	return m.errs
}

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
