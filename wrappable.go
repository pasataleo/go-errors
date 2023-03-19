package errors

type Unwrappable interface {
	error
	Unwrap() error
}

func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if unwrappable, ok := err.(Unwrappable); ok {
		return unwrappable.Unwrap()
	}
	return nil
}
