package endpoint

import (
	"net/http"
)

func Pinger(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("pong"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
