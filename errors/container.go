package errors

type DataContainer[Data any] interface {
	error
	GetEmbeddedData() Data
}

func GetEmbeddedData[Data any](err error) (Data, bool) {
	var data Data
	if err == nil {
		return data, false
	}

	if container, ok := err.(DataContainer[Data]); ok {
		return container.GetEmbeddedData(), true
	}

	return data, false
}
