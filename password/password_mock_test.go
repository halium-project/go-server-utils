package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashManagerMock_Impl(t *testing.T) {
	assert.Implements(t, (*HashManager)(nil), new(HashManagerMock))
}

func Test_HashManagerMock_HashWithSalt(t *testing.T) {
	mock := new(HashManagerMock)

	mock.On("HashWithSalt", "some-password").Return("some-hash", "some-salt", nil).Once()
	hash, salt, err := mock.HashWithSalt("some-password")

	assert.Equal(t, "some-hash", hash)
	assert.Equal(t, "some-salt", salt)
	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func Test_HashManagerMock_Hash(t *testing.T) {
	mock := new(HashManagerMock)

	mock.On("Hash", "some-password").Return("some-hash", nil).Once()
	hash, err := mock.Hash("some-password")

	assert.Equal(t, "some-hash", hash)
	assert.NoError(t, err)

	mock.AssertExpectations(t)
}

func Test_HashManagerMock_ValidateWithSalt(t *testing.T) {
	mock := new(HashManagerMock)

	mock.On("ValidateWithSalt", "some-password", "some-salt", "some-hash").Return(true, nil).Once()

	identical, err := mock.ValidateWithSalt("some-password", "some-salt", "some-hash")
	assert.NoError(t, err)
	assert.True(t, identical)

	mock.AssertExpectations(t)
}

func Test_HashManagerMock_Validate(t *testing.T) {
	mock := new(HashManagerMock)

	mock.On("Validate", "some-password", "some-hash").Return(true, nil).Once()

	identical, err := mock.Validate("some-password", "some-hash")
	assert.NoError(t, err)
	assert.True(t, identical)

	mock.AssertExpectations(t)
}
