package uuid

import "github.com/stretchr/testify/mock"

// ProducerMock is a mock implementation of the `Manager` interface.
type ProducerMock struct {
	mock.Mock
}

// New is an implementation of the `Manager` interface.
func (t *ProducerMock) New() string {
	return t.Called().String(0)
}

func (t *ProducerMock) IsValid(input string) bool {
	return t.Called(input).Bool(0)
}
