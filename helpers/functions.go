package helpers

import (
	"strings"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Syntactic sugar for JS -> startsWith()
func (s String) StartsWith(toMatch string) bool {
	return strings.HasPrefix(string(s), toMatch)
}

// Syntactic sugar for JS -> endsWith()
func (s String) EndsWith(toMatch string) bool {
	return strings.HasSuffix(string(s), toMatch)
}
