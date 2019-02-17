package endpoint

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Pinger(t *testing.T) {
	r := httptest.NewRequest("GET", "http://example.com/ping", nil)
	w := httptest.NewRecorder()

	Pinger(w, r)

	res := w.Result()
	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "pong", string(body))
}
