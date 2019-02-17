package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	err := New(Internal, "some-error")

	assert.JSONEq(t, `{
		"kind": "internalError",
		"message": "some-error"
	}`, string(err.IntoJSON()))
}

func Test_Errorf(t *testing.T) {
	err := Errorf(Internal, "some-error: %d", 42)

	assert.JSONEq(t, `{
		"kind": "internalError",
		"message": "some-error: 42"
	}`, string(err.IntoJSON()))
}

func Test_Wrap(t *testing.T) {
	err := Wrap(New(NotFound, "some-parent-error"), "some-details")

	assert.JSONEq(t, `{
		"kind": "notFound",
		"message": "some-details",
		"reason": {
			"kind": "notFound",
			"message": "some-parent-error"
		}
	}`, string(err.IntoJSON()))
}

func Test_Wrapf(t *testing.T) {
	err := Wrapf(New(Internal, "some-parent-error"), "some-details: %d", 42)

	assert.JSONEq(t, `{
		"kind": "internalError",
		"message": "some-details: 42",
		"reason": {
			"kind": "internalError",
			"message": "some-parent-error"
		}
	}`, string(err.IntoJSON()))
}

func Test_Wrap_system_error(t *testing.T) {
	err := Wrap(fmt.Errorf("some-parent-error"), "some-details")

	assert.JSONEq(t, `{
		"kind": "internalError",
		"message": "some-details",
		"reason": {
			"kind": "internalError",
			"message": "some-parent-error"
		}
	}`, string(err.IntoJSON()))
}

func Test_Wrap_internal(t *testing.T) {
	internalError := Wrap(New(NotFound, "some-error"), "some-details-1")

	err := Wrap(internalError, "some-details-2")

	assert.JSONEq(t, `{
		"kind": "notFound",
		"message": "some-details-2",
		"reason": {
			"kind": "notFound",
			"message": "some-details-1",
			"reason": {
				"kind": "notFound",
				"message": "some-error"
			}
		}
	}`, string(err.IntoJSON()))
}

func Test_Error_Error(t *testing.T) {
	err := New(Internal, "some-error")

	assert.JSONEq(t, `{
		"kind": "internalError",
		"message": "some-error"
	}`, err.Error())
}

func Test_IsKind(t *testing.T) {
	err := New(Internal, "some-error")

	assert.True(t, IsKind(err, Internal))
	assert.False(t, IsKind(err, Validation))
	assert.False(t, IsKind(fmt.Errorf("some-error"), Validation))
	assert.False(t, IsKind(nil, Validation))
}
