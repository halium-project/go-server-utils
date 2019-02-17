package uuid

import uuid "github.com/satori/go.uuid"

// GoUUID is an implementation of `Manager` with the 'go.uuid' lib.
type GoUUID struct{}

// NewGoUUID instantiate a new GoUUID.
func NewGoUUID() *GoUUID {
	return &GoUUID{}
}

// New is an implementation of the `Manager` interface.
func (t *GoUUID) New() string {
	return uuid.NewV4().String()
}

func (t *GoUUID) IsValid(input string) bool {
	_, err := uuid.FromString(input)
	return err == nil
}
