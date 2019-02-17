package is

import (
	"errors"
)

func NumberPositif(number int) error {
	if number < 0 {
		return errors.New(UnexpectedValue)
	}

	return nil
}
