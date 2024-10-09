package errors

// DataContainer is an interface for errors that contain embedded data.
type DataContainer interface {
	error
	EmbeddedData() map[string]interface{}
}

// GetAllEmbeddedData returns all embedded data from an error if it exists. This function will automatically unwrap
// errors that are wrapped in other errors to find all metadata associated with the error.
func GetAllEmbeddedData(err error) map[string]interface{} {
	if err == nil {
		return nil
	}

	metadata := make(map[string]interface{})

	if container, ok := err.(DataContainer); ok {
		for key, value := range container.EmbeddedData() {
			metadata[key] = value
		}
	}

	// Recursively unwrap errors to find all metadata.
	wrapped := GetAllEmbeddedData(Unwrap(err))
	for key, value := range wrapped {
		metadata[key] = value
	}

	return metadata
}

// GetEmbeddedData returns the embedded data from an error if it exists. This function will automatically unwrap
// errors that are wrapped in other errors to find all metadata associated with the error.
func GetEmbeddedData[Data any](err error, key string) (Data, bool) {
	var data Data
	if err == nil {
		return data, false
	}

	if container, ok := err.(DataContainer); ok {
		data := container.EmbeddedData()
		if value, ok := data[key]; ok {
			if d, ok := value.(Data); ok {
				return d, true
			}
		}
	}

	// We didn't find the data in the current error, so let's try unwrapping it.
	return GetEmbeddedData[Data](Unwrap(err), key)
}

// GetEmbeddedDataUnsafe returns the embedded data from an error if it exists. This function will automatically unwrap
// errors that are wrapped in other errors to find all metadata associated with the error.
func GetEmbeddedDataUnsafe(err error, key string) interface{} {
	if err == nil {
		return nil
	}

	if container, ok := err.(DataContainer); ok {
		data := container.EmbeddedData()
		if value, ok := data[key]; ok {
			return value
		}
	}

	// We didn't find the data in the current error, so let's try unwrapping it.
	return GetEmbeddedDataUnsafe(Unwrap(err), key)
}
