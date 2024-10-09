package errors

// Unwrappable is an interface for errors that can be unwrapped.
type Unwrappable interface {
	error
	Unwrap() error
}

// Unwrap returns the unwrapped error if it exists.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if unwrappable, ok := err.(Unwrappable); ok {
		return unwrappable.Unwrap()
	}
	return nil
}
