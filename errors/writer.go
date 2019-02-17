package errors

import (
	"log"
	"net/http"

	"github.com/halium-project/go-server-utils/env"
)

func IntoResponse(w http.ResponseWriter, err error) {
	switch err.(type) {
	case *Error:
		WriteError(w, err.(*Error))

	default:
		log.Printf("unhandled error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}

func WriteError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/json")

	switch err.Kind {
	case Internal:
		w.WriteHeader(http.StatusInternalServerError)

		if env.MustGetEnv("ENV") == "production" {
			// Don't write the body of internal errors in production
			return
		}

	case NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)

	case NotFound:
		w.WriteHeader(http.StatusNotFound)

	case Validation:
		w.WriteHeader(http.StatusUnprocessableEntity)

	case BadRequest:
		w.WriteHeader(http.StatusBadRequest)

	case Forbidden:
		w.WriteHeader(http.StatusForbidden)

	default:
		log.Printf("unknown error kind: %s -> %s", err.Kind, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, writeErr := w.Write([]byte(err.Error()))
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
