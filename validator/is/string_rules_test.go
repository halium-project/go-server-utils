package is

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newStringRule(t *testing.T) {
	foobarValidator := func(in string) bool {
		return in == "foobar"
	}

	rule := newStringRule(foobarValidator)
	err := rule("foobar")
	assert.NoError(t, err)

	err = rule("john doe")
	assert.EqualError(t, err, InvalidFormat)
}

func Test_InRange(t *testing.T) {
	validator := StringInRange(3, 5)

	assert.EqualError(t, validator(""), TooShort)
	assert.EqualError(t, validator("this is too long"), TooLong)
	assert.NoError(t, validator("1234"))
}

func Test_OnOfString(t *testing.T) {
	validator := OnOfString("foo", "bar")

	assert.EqualError(t, validator("baz"), UnexpectedValue)
	assert.EqualError(t, validator(""), UnexpectedValue)
	assert.NoError(t, validator("foo"))
	assert.NoError(t, validator("bar"))
}

func Test_MatchingString(t *testing.T) {
	validator := MatchingString("foo.+baz")

	assert.EqualError(t, validator("baz"), InvalidFormat)
	assert.EqualError(t, validator(""), InvalidFormat)
	assert.NoError(t, validator("foobarbaz"))
}
