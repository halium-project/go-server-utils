package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProducerMockMock_Impl(t *testing.T) {
	assert.Implements(t, (*Producer)(nil), new(ProducerMock))
}

func Test_ProducerMock_New(t *testing.T) {
	mock := new(ProducerMock)

	mock.On("New").Return("some-mock-uuid").Once()
	id := mock.New()
	assert.Equal(t, "some-mock-uuid", id)

	mock.AssertExpectations(t)
}

func Test_ProducerMock_IsValid(t *testing.T) {
	mock := new(ProducerMock)

	mock.On("IsValid", "some-input-to-check").Return(true).Once()
	isValid := mock.IsValid("some-input-to-check")
	assert.True(t, isValid)

	mock.AssertExpectations(t)
}
