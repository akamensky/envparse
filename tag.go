package envparse

import (
	"errors"
	"regexp"
	"strings"
)

type tagType struct {
	name         string
	required     bool
	defaultValue string
}

func parseTag(t string) (*tagType, error) {
	name := ""
	required := false
	defaultValue := ""

	tagParts := strings.Split(t, ",")
	for _, tagPart := range tagParts {
		if tagPart == "required" {
			required = true
		} else if strings.HasPrefix(tagPart, "default=") {
			defaultValue = strings.TrimPrefix(tagPart, "default=")
		} else if strings.HasPrefix(tagPart, "name=") {
			name = strings.ToUpper(regexp.MustCompile(`[^A-Za-z0-9]`).ReplaceAllString(strings.TrimPrefix(tagPart, "name="), ""))
		}
	}

	if name == "" {
		return nil, errors.New("field tag must provide a name for field")
	}
	if required && defaultValue != "" {
		return nil, errors.New("field tag cannot be required and have default value at the same time")
	}

	result := &tagType{
		name:         name,
		required:     required,
		defaultValue: defaultValue,
	}

	return result, nil
}