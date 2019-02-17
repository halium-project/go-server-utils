package password

import "github.com/stretchr/testify/mock"

type HashManagerMock struct {
	mock.Mock
}

func (t *HashManagerMock) HashWithSalt(password string) (string, string, error) {
	args := t.Called(password)
	return args.String(0), args.String(1), args.Error(2)
}

func (t *HashManagerMock) Hash(password string) (string, error) {
	args := t.Called(password)
	return args.String(0), args.Error(1)
}

func (t *HashManagerMock) Validate(password string, hash string) (bool, error) {
	args := t.Called(password, hash)
	return args.Bool(0), args.Error(1)
}

func (t *HashManagerMock) ValidateWithSalt(password string, salt string, hash string) (bool, error) {
	args := t.Called(password, salt, hash)
	return args.Bool(0), args.Error(1)
}
