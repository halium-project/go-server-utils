package is

import (
	"fmt"
	"reflect"
)

func ArrayInRange(min int, max int) func(in interface{}) error {
	return func(in interface{}) error {
		if in == nil {
			return fmt.Errorf(MissingField)
		}

		rv := reflect.ValueOf(in)

		if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
			return fmt.Errorf(InvalidType)
		}

		if rv.Len() < min {
			return fmt.Errorf(TooShort)
		}

		if rv.Len() > max {
			return fmt.Errorf(TooLong)
		}

		return nil
	}
}
