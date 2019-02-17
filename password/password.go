package password

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type HashManager interface {
	Hash(password string) (string, error)
	HashWithSalt(password string) (string, string, error)
	Validate(password string, hash string) (bool, error)
	ValidateWithSalt(password string, salt string, hash string) (bool, error)
}

type Hasher struct{}

func NewPasswordHasher() *Hasher {
	return &Hasher{}
}

func (t *Hasher) HashWithSalt(password string) (string, string, error) {
	if password == "" {
		return "", "", errors.New("empty password")
	}

	salt := uuid.NewV4().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate password")
	}

	return string(hashedPassword), salt, nil
}

func (t *Hasher) Hash(password string) (string, error) {
	if password == "" {
		return "", errors.New("empty password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate password")
	}

	return string(hashedPassword), nil
}

func (t *Hasher) Validate(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	if err != nil {
		return false, errors.Wrap(err, "failed to generate password")
	}

	return true, nil
}

func (t *Hasher) ValidateWithSalt(password string, salt string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	if err != nil {
		return false, errors.Wrap(err, "failed to generate password")
	}

	return true, nil
}
