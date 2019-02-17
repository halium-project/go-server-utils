package is

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Array(t *testing.T) {
	var tests = []struct {
		Title     string
		Validator func(interface{}) error
		In        interface{}
		Out       string
	}{
		{
			Title:     "range with nil",
			Validator: ArrayInRange(3, 5),
			In:        nil,
			Out:       MissingField,
		},
		{
			Title:     "range with too few elements",
			Validator: ArrayInRange(3, 5),
			In:        []int{1, 2},
			Out:       TooShort,
		},
		{
			Title:     "range with in invalid type",
			Validator: ArrayInRange(3, 5),
			In:        "string",
			Out:       InvalidType,
		},
		{
			Title:     "range with too many elements",
			Validator: ArrayInRange(3, 5),
			In:        []int{1, 2, 3, 4, 5, 6, 7},
			Out:       TooLong,
		},
		{
			Title:     "range with valid value",
			Validator: ArrayInRange(3, 5),
			In:        []int{1, 2, 3, 4},
			Out:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(tt *testing.T) {
			err := test.Validator(test.In)
			if test.Out == "" {
				assert.NoError(tt, err)
			} else {
				assert.EqualError(tt, err, test.Out)
			}
		})
	}
}
