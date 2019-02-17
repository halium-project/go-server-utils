package validator

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/halium-project/go-server-utils/errors"
	"github.com/halium-project/go-server-utils/validator/is"
)

type numberValidator func(int) error
type stringValidator func(string) error
type arrayValidator func(interface{}) error

type Validator struct {
	inner *errors.ValidationError
}

func New() *Validator {
	return &Validator{
		inner: errors.NewValidationError(),
	}
}

func (t *Validator) CheckString(fieldName string, in string, presence is.Presence, validators ...stringValidator) *Validator {
	var err error

	if len(in) == 0 {
		if presence == is.Required {
			t.inner.AddError(fieldName, is.MissingField)
		}

		return t
	}

	for _, validator := range validators {
		err = validator(in)
		if err != nil {
			t.inner.AddError(fieldName, err.Error())
		}
	}

	return t
}

func (t *Validator) CheckNumber(fieldName string, in int, presence is.Presence, validators ...numberValidator) *Validator {
	var err error

	for _, validator := range validators {
		err = validator(in)
		if err != nil {
			t.inner.AddError(fieldName, err.Error())
		}
	}

	return t
}

func (t *Validator) CheckArray(fieldName string, in interface{}, validators ...arrayValidator) *Validator {
	var err error

	for _, validator := range validators {
		err = validator(in)
		if err != nil {
			t.inner.AddError(fieldName, err.Error())
		}
	}

	return t
}

type StructValidation func(*Validator) *Validator

func NoValidation(t *Validator) *Validator {
	return t
}

func (t *Validator) CheckStruct(fieldName string, in interface{}, presence is.Presence, structValidations StructValidation) *Validator {
	if in == nil || reflect.ValueOf(in).IsNil() {
		if presence == is.Required {
			t.inner.AddError(fieldName, is.MissingField)
		}

		return t
	}

	return structValidations(t)
}

func (t *Validator) CustomStructValidation(validation StructValidation) *Validator {
	return validation(t)
}

func (t *Validator) CheckEachString(fieldName string, fields []string, validators ...stringValidator) *Validator {
	var err error

	for i, field := range fields {
		for _, validator := range validators {
			err = validator(field)
			if err != nil {
				t.inner.AddError(
					strings.Join([]string{fieldName, "[", strconv.Itoa(i), "]"}, ""),
					err.Error(),
				)
				break
			}
		}
	}

	return t
}

func (t *Validator) Run() error {
	if len(t.inner.Errors) > 0 {
		return t.inner.IntoError()
	}

	return nil
}
