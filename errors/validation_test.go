package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_AddValidationError(t *testing.T) {
	err := NewValidationError().
		AddError("some-resource-1", "some-code-1").
		AddError("some-resource-2", "some-code-2")

	assert.JSONEq(t, string(err.IntoError().IntoJSON()), `{
		"kind": "validationError",
		"errors": {
			"some-resource-1": "some-code-1",
			"some-resource-2": "some-code-2"
		}
	}`)
}
