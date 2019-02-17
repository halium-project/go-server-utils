package validator

import (
	"testing"

	"github.com/halium-project/go-server-utils/validator/is"
	"github.com/stretchr/testify/assert"
)

func Test_CheckString(t *testing.T) {
	value := "invalid-value"

	err := New().
		CheckString("foobar", value, is.Required, is.Email).
		Run()

	assert.JSONEq(t, `{
		"kind":"validationError",
		"errors": {
			"foobar": "INVALID_FORMAT"
		}
	}`, err.Error())
}

func Test_CheckString_with_empty_field(t *testing.T) {
	emptyValue := ""

	err := New().
		CheckString("required-field", emptyValue, is.Required, is.Email).
		CheckString("optional-field", emptyValue, is.Optional).
		Run()

	assert.JSONEq(t, `{
		"kind":"validationError",
		"errors": {
			"required-field": "MISSING_FIELD"
		}
	}`, err.Error())
}

func Test_CheckString_with_array(t *testing.T) {
	emailList := []string{
		"valid@email.fr",
		"foobar",
		"",
	}

	err := New().
		CheckEachString("foobar", emailList, is.Email).
		Run()

	assert.JSONEq(t, `{
		"kind": "validationError",
		"errors": {
			"foobar[1]": "INVALID_FORMAT",
			"foobar[2]": "INVALID_FORMAT"
		}
	}`, err.Error())
}

func Test_CheckString_with_no_error(t *testing.T) {
	value := "valid@email.fr"

	err := New().
		CheckString("foobar", value, true, is.Email).
		Run()

	assert.NoError(t, err)
}

func Test_CheckArray_with_no_error(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	err := New().
		CheckArray("field-name", values, is.ArrayInRange(3, 10)).
		Run()

	assert.NoError(t, err)
}

func Test_CheckArray_with_range_error(t *testing.T) {
	values := []int{}

	err := New().
		CheckArray("field-name", values, is.ArrayInRange(3, 10)).
		Run()

	assert.JSONEq(t, `{
		"kind":"validationError",
		"errors":{
			"field-name": "TOO_SHORT"
		}
	}`, err.Error())
}

func Test_CheckStruct(t *testing.T) {
	type keyStore struct {
		Key string
	}

	someInvalidValue := keyStore{Key: "not an email"}

	err := New().
		CheckStruct("structName", &someInvalidValue, is.Required, func(t *Validator) *Validator {
			return t.CheckString("structName.key", someInvalidValue.Key, is.Required, is.Email)
		}).
		Run()

	assert.JSONEq(t, `{
		"kind":"validationError",
		"errors":{
			"structName.key": "INVALID_FORMAT"
		}
	}`, err.Error())
}

func Test_CheckStruct_with_nil_and_optional_struct(t *testing.T) {
	type keyStore struct {
		Key string
	}

	// is nil
	var someInvalidValue *keyStore

	err := New().
		CheckStruct("structName", someInvalidValue, is.Optional, func(t *Validator) *Validator {
			return t.CheckString("structName.key", someInvalidValue.Key, is.Required, is.Email)
		}).
		Run()

	assert.NoError(t, err)
}

func Test_CheckStruct_optional_with_nil(t *testing.T) {
	err := New().
		CheckStruct("field-name", nil, is.Optional, NoValidation).
		Run()

	assert.NoError(t, err)
}

func Test_CheckStruct_required_with_nil(t *testing.T) {
	err := New().
		CheckStruct("field-name", nil, is.Required, NoValidation).
		Run()

	assert.JSONEq(t, `{
		"kind":"validationError",
		"errors":{
			"field-name": "MISSING_FIELD"
		}
	}`, err.Error())
}
