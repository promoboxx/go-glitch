package validate

import "fmt"

// ValidationError is an interface for returnign invalid request data
type ValidationError interface {
	GetInvalidFields() map[string]string
	AddInvalidField(field string, reason string)
	Error() string
	String() string
}

type validationError struct {
	invalidFields map[string]string
}

// NewValidationError returns a new ValidationError. Implemented by validationError
func NewValidationError(field string, reason string) ValidationError {
	ve := &validationError{invalidFields: map[string]string{}}
	ve.AddInvalidField(field, reason)
	return ve
}

func InitValidationError() ValidationError {
	return &validationError{invalidFields: map[string]string{}}
}

// AddInvalidField adds a field to the map of all invalid fields that failed validation
func (ve *validationError) AddInvalidField(field string, reason string) {
	ve.invalidFields[field] = reason
}

// GetInvalidFields returns the map of invalid field names to reasons
func (ve *validationError) GetInvalidFields() map[string]string {
	return ve.invalidFields
}

// Error allows to implement the error interface
func (ve *validationError) Error() string {
	builder := "Invalid fields-"
	for k, v := range ve.invalidFields {
		builder += fmt.Sprintf(" %s: %s,", k, v)
	}

	builder = builder[:len(builder)-1]

	return builder
}

// String will give back a formatted string of invalid field names to reasons
// ex:
// "Invalid fields- email: was not formatted properly, phone: not enough numbers"
func (ve *validationError) String() string {
	return ve.Error()
}
