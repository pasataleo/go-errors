package errors

type DataContainer interface {
	error
	GetEmbeddedData() interface{}
}

func GetEmbeddedData[Data any](err error) (Data, bool) {
	var data Data
	if err == nil {
		return data, false
	}

	if container, ok := err.(DataContainer); ok {
		if value, ok := container.GetEmbeddedData().(Data); ok {
			return value, true
		}
	}

	return data, false
}
