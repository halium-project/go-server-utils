package errors

import (
	"log"
	"net/http"
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

	case NotAuthorized:
		w.WriteHeader(http.StatusUnauthorized)

	case NotFound:
		w.WriteHeader(http.StatusNotFound)

	case Validation:
		w.WriteHeader(http.StatusUnprocessableEntity)

	case InvalidJSON:
		w.WriteHeader(http.StatusBadRequest)

	case BadRequest:
		w.WriteHeader(http.StatusBadRequest)

	case InvalidMethod:
		w.WriteHeader(http.StatusMethodNotAllowed)

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
