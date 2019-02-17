package response

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WriteOK(t *testing.T) {
	w := httptest.NewRecorder()

	type response struct {
		Key string `json:"key"`
	}

	Write(w, http.StatusOK, &response{Key: "value"})

	res := w.Result()

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	assert.JSONEq(t, `{
		"key": "value"
	}`, string(body))
}
