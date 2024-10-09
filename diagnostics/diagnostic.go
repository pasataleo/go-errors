package diagnostics

import "github.com/pasataleo/go-errors/errors"

// Diagnostic represents a diagnostic message. This is an advanced kind of error that can be used to provide additional
// information about a problem. It can be used to provide more context about an error, warning, or information message.
type Diagnostic interface {

	// Severity returns the severity of the diagnostic.
	Severity() Severity

	// Summary returns a short summary of the diagnostic.
	Summary() string

	// Detail returns a detailed description of the diagnostic. This should be considered optional.
	Detail() string

	// Metadata returns additional metadata about the diagnostic. This can be used to provide additional context about the
	// diagnostic.
	Metadata() map[string]interface{}

	// Error returns the diagnostic as an error. The provided severity is used to determine which level of severity is
	// considered an error. For example, if the severity is SeverityError, then only errors should return an error here
	// with SeverityInfo and SeverityWarning returning nil.
	Error(severity Severity) error
}

// diagnostic is the default implementation of the Diagnostic interface.
type diagnostic struct {
	severity Severity
	summary  string
	detail   string
	metadata map[string]interface{}
}

// Severity implements the Diagnostic interface.
func (d *diagnostic) Severity() Severity {
	return d.severity
}

// Summary implements the Diagnostic interface.
func (d *diagnostic) Summary() string {
	return d.summary
}

// Detail implements the Diagnostic interface.
func (d *diagnostic) Detail() string {
	return d.detail
}

// Metadata implements the Diagnostic interface.
func (d *diagnostic) Metadata() map[string]interface{} {
	return d.metadata
}

// Error implements the Diagnostic interface.
func (d *diagnostic) Error(incl Severity) error {
	if d.severity < incl {
		return nil
	}

	msg := d.summary
	if d.detail != "" {
		msg += ": " + d.detail
	}

	errorCode := errors.ErrorCodeUnknown
	if ec, ok := d.metadata["error_code"].(errors.ErrorCode); ok {
		errorCode = ec
	}

	err := errors.Newf(nil, errorCode, "[%s] %s", d.severity, msg)
	for key, value := range d.metadata {
		if key == "error_code" {
			continue
		}
		err = errors.Embed(err, key, value)
	}
	return err
}

// DiagnosticBuilder is a builder for creating diagnostics.
type DiagnosticBuilder struct {
	diagnostic diagnostic
}

// Info creates a new diagnostic with the info severity.
func Info(summary string) *DiagnosticBuilder {
	return &DiagnosticBuilder{
		diagnostic: diagnostic{
			severity: SeverityInfo,
			summary:  summary,
			metadata: make(map[string]interface{}),
		},
	}
}

// Warning creates a new diagnostic with the warning severity.
func Warning(summary string) *DiagnosticBuilder {
	return &DiagnosticBuilder{
		diagnostic: diagnostic{
			severity: SeverityWarning,
			summary:  summary,
		},
	}
}

// Error creates a new diagnostic with the error severity.
func Error(summary string) *DiagnosticBuilder {
	return &DiagnosticBuilder{
		diagnostic: diagnostic{
			severity: SeverityError,
			summary:  summary,
		},
	}
}

// Detail sets the detail of the diagnostic.
func (b *DiagnosticBuilder) Detail(detail string) *DiagnosticBuilder {
	b.diagnostic.detail = detail
	return b
}

// AddAllMetadata adds multiple metadata entries to the diagnostic.
func (b *DiagnosticBuilder) AddAllMetadata(metadata map[string]interface{}) *DiagnosticBuilder {
	for key, value := range metadata {
		b.diagnostic.metadata[key] = value
	}
	return b
}

// AddMetadata adds metadata to the diagnostic.
func (b *DiagnosticBuilder) AddMetadata(key string, value interface{}) *DiagnosticBuilder {
	b.diagnostic.metadata[key] = value
	return b
}

// SetErrorCode sets the error code of the diagnostic via the metadata.
func (b *DiagnosticBuilder) SetErrorCode(code errors.ErrorCode) *DiagnosticBuilder {
	b.diagnostic.metadata["error_code"] = code
	return b
}

// Build creates the diagnostic.
func (b *DiagnosticBuilder) Build() Diagnostic {
	return &b.diagnostic
}
