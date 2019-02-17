package errors

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_IntoResponse_with_InternalError(t *testing.T) {
	w := httptest.NewRecorder()

	err := os.Setenv("ENV", "not-production")
	require.NoError(t, err)

	IntoResponse(w, New(Internal, "some-error"))

	res, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{
		"kind": "internalError",
		"message": "some-error"
	}`, string(res))
}

func Test_IntoResponse_with_InternalError_and_prod(t *testing.T) {
	w := httptest.NewRecorder()

	err := os.Setenv("ENV", "production")
	require.NoError(t, err)

	IntoResponse(w, New(Internal, "some-error"))

	res, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Empty(t, string(res))
}

func Test_IntoResponse_with_ValidationError(t *testing.T) {
	w := httptest.NewRecorder()

	err := os.Setenv("ENV", "production")
	require.NoError(t, err)

	IntoResponse(w, NewValidationError().AddError("some-target", "some-error").IntoError())

	res, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.JSONEq(t, `{
		"kind": "validationError",
		"errors": {
			"some-target": "some-error"
		}
	}`, string(res))

}
