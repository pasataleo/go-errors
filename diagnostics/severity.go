package diagnostics

type Severity int

const (
	SeverityUnknown Severity = iota
	SeverityInfo
	SeverityWarning
	SeverityError
)

// Merge returns the highest severity of the two.
func (s Severity) Merge(other Severity) Severity {
	if s > other {
		return s
	}
	return other
}

// MergeSeverities returns the highest severity of the given severities.
func MergeSeverities(severities ...Severity) Severity {
	var m Severity
	for _, s := range severities {
		m = m.Merge(s)
	}
	return m
}

func (s Severity) String() string {
	switch s {
	case SeverityInfo:
		return "info"
	case SeverityWarning:
		return "warning"
	case SeverityError:
		return "error"
	default:
		return "unknown"
	}
}
