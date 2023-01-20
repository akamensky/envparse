package envparse

import (
	"regexp"
	"strings"
)

func SetPrefix(s string) {
	defaultPrefix = strings.ToUpper(regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(s, ""))
}

func SetMaxDepth(i int) {
	defaultMaxDepth = i
}
