package errors

import (
	"encoding/json"
	"fmt"
	"io"
)

type ErrorKind string

const (
	BadRequest         ErrorKind = "badRequest"
	Forbidden          ErrorKind = "forbidden"
	Internal           ErrorKind = "internalError"
	InvalidCredentials ErrorKind = "invalidCredentials"
	NotAuthorized      ErrorKind = "notAuthorized"
	NotFound           ErrorKind = "notFound"
	Validation         ErrorKind = "validationError"
)

type Error struct {
	Kind    ErrorKind         `json:"kind"`
	Message string            `json:"message,omitempty"`
	Reason  *Error            `json:"reason,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func New(kind ErrorKind, msg string) *Error {
	return &Error{
		Kind:    kind,
		Message: msg,
		Reason:  nil,
	}
}

func IsKind(err error, kind ErrorKind) bool {
	typedError, ok := err.(*Error)
	if !ok {
		return false
	}

	return typedError.Kind == kind
}

func IsUnexpected(err error) bool {
	if err == nil {
		return false
	}

	typedError, ok := err.(*Error)
	if !ok {
		return true
	}

	return typedError.Kind == Internal
}

func FromReader(in io.Reader) *Error {
	var res Error

	if in == nil {
		return New(Internal, "reader empty")
	}

	err := json.NewDecoder(in).Decode(&res)
	if err != nil {
		return Wrap(err, "failed to decode reader")
	}

	return &res
}

func Errorf(kind ErrorKind, format string, v ...interface{}) *Error {
	return New(kind, fmt.Sprintf(format, v...))
}

func Wrap(err error, msg string) *Error {
	var (
		internalError *Error
		ok            bool
	)

	internalError, ok = err.(*Error)
	if !ok {
		// If err is not and Error type, generate an InternalError by default
		internalError = New(Internal, err.Error())
	}

	return &Error{
		Kind:    internalError.Kind,
		Message: msg,
		Reason:  internalError,
	}
}

func Wrapf(err error, format string, v ...interface{}) *Error {
	return Wrap(err, fmt.Sprintf(format, v...))
}

func (t *Error) IntoJSON() []byte {
	res, err := json.Marshal(t)
	if err != nil {
		panic(fmt.Sprintf("fail to marshal error: %s", err))
	}
	return res
}

func (t *Error) Error() string {
	return string(t.IntoJSON())
}
