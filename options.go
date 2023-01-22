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

// SetUnsetEnv sets a flag that if `true` will ensure that all environment variables using defined prefix
// will not be readable by the process or by the children of the process.
//
// NOTE: the original environment variable is still available to be read from `/proc/<pid>/environ` or `ps eww <pid>`
// as there are no system mechanisms available to overwrite those.
func SetUnsetEnv(unset bool) {
	unsetEnv = unset
}
