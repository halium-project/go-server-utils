package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GoUUID_impl_Producer(t *testing.T) {
	assert.Implements(t, (*Producer)(nil), new(GoUUID))
}

func Test_GoUUID_New_and_IsValid(t *testing.T) {
	uuidProducer := NewGoUUID()

	id := uuidProducer.New()

	assert.True(t, uuidProducer.IsValid(id))
}
