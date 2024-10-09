package diagnostics

import "github.com/pasataleo/go-errors/errors"

type Diagnostics []Diagnostic

// Match returns true if any of the diagnostics match either the provider severity or one of higher precedence. For
// example diags.Match(SeverityWarning) will return true if any of the diagnostics are SeverityWarning or SeverityError.
func (d Diagnostics) Match(severity Severity) bool {
	for _, diagnostic := range d {
		if diagnostic.Severity() >= severity {
			return true
		}
	}
	return false
}

// Filter returns a new Diagnostics instance with only the diagnostics that match the provided severity.
func (d Diagnostics) Filter(severity Severity) Diagnostics {
	var diagnostics Diagnostics
	for _, diagnostic := range d {
		if diagnostic.Severity() == severity {
			diagnostics = append(diagnostics, diagnostic)
		}
	}
	return diagnostics
}

// Severity returns the highest severity of all the diagnostics.
func (d Diagnostics) Severity() Severity {
	var m Severity
	for _, diagnostic := range d {
		if sev := diagnostic.Severity(); sev > m {
			m = sev
		}
	}
	return m
}

// Error returns an error if any of the diagnostics are of the provided severity or higher.
func (d Diagnostics) Error(incl Severity) error {
	var err error
	for _, diagnostic := range d {
		if e := diagnostic.Error(incl); e != nil {
			err = errors.Append(err, e)
		}
	}
	return err
}
