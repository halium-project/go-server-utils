package errors

type ValidationError Error

type ValidationDetail struct {
	Field string `json:"field"`
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		Kind:   Validation,
		Reason: nil,
		Errors: map[string]string{},
	}
}

func (t *ValidationError) AddError(field string, code string) *ValidationError {
	t.Errors[field] = code

	return t
}

func (t *ValidationError) IntoError() *Error {
	res := Error(*t)
	return &res
}
