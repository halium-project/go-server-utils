package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PasswordHasher(t *testing.T) {
	hasher := NewPasswordHasher()

	hash, err := hasher.Hash("some-big-password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	identical, err := hasher.Validate("some-big-password", hash)
	assert.NoError(t, err)
	assert.True(t, identical)
}

func Test_PasswordHasher_with_salt(t *testing.T) {
	hasher := NewPasswordHasher()

	hash, salt, err := hasher.HashWithSalt("some-big-password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, salt)

	identical, err := hasher.ValidateWithSalt("some-big-password", salt, hash)
	assert.NoError(t, err)
	assert.True(t, identical)
}

func Test_PasswordHasher_HashWithSalt_len_salt(t *testing.T) {
	hasher := NewPasswordHasher()

	_, salt, err := hasher.HashWithSalt("some-big-password")
	assert.NoError(t, err)

	assert.True(t, len(salt) > 8)
}

func Test_PasswordHasher_Hash_empty_password(t *testing.T) {
	hasher := NewPasswordHasher()

	hash, err := hasher.Hash("")
	assert.Empty(t, hash)
	assert.EqualError(t, err, "empty password")
}

func Test_PasswordHasher_HashWithSalt_empty_password(t *testing.T) {
	hasher := NewPasswordHasher()

	hash, salt, err := hasher.HashWithSalt("")
	assert.Empty(t, hash)
	assert.Empty(t, salt)
	assert.EqualError(t, err, "empty password")
}

func Test_PasswordHasher_Validate_bad_hash_encoding(t *testing.T) {
	hasher := NewPasswordHasher()

	identical, err := hasher.Validate("some-big-password", "no-hexa-code")
	assert.False(t, identical)
	assert.EqualError(t, err, "failed to generate password: crypto/bcrypt: hashedSecret too short to be a bcrypted password")
}

func Test_PasswordHasher_ValidateWithSalt_bad_hash_encoding(t *testing.T) {
	hasher := NewPasswordHasher()

	identical, err := hasher.ValidateWithSalt("some-big-password", "some-salt", "no-hexa-code")
	assert.False(t, identical)
	assert.EqualError(t, err, "failed to generate password: crypto/bcrypt: hashedSecret too short to be a bcrypted password")
}

func Test_PasswordHasher_Validate_not_invalid(t *testing.T) {
	hasher := NewPasswordHasher()

	hash, err := hasher.Hash("some-big-password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	identical, err := hasher.Validate("not-identical", hash)
	assert.False(t, identical)
	assert.NoError(t, err)
}

func Test_PasswordHasher_ValidateWithSalt_not_invalid(t *testing.T) {
	hasher := NewPasswordHasher()

	hash, salt, err := hasher.HashWithSalt("some-big-password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, salt)

	identical, err := hasher.ValidateWithSalt("not-identical", salt, hash)
	assert.False(t, identical)
	assert.NoError(t, err)
}
