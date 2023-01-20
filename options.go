package envparse

import (
	"regexp"
	"strings"
)

// SetPrefix sets new prefix to use for scanning. Default is `APP`
func SetPrefix(s string) {
	defaultPrefix = strings.ToUpper(regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(s, ""))
}

// SetMaxDepth sets maximum embedded depth allowed for structs and slices.
// Crossing this limit will cause an error. Serves as a fail-safe against accidental infinite loops.
// Default is 100
func SetMaxDepth(i int) {
	defaultMaxDepth = i
}
