package env

import (
	"fmt"
	"os"
)

func MustGetEnv(key string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		panic(fmt.Sprintf("missing %q env variable", key))
	}

	return value
}
