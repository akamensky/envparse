package envparse

import "fmt"

type ErrorList struct {
	errors []error
}

func (e *ErrorList) Error() string {
	return fmt.Sprintf("%d errors happened during parsing environment variables", len(e.errors))
}

func (e *ErrorList) FullError() string {
	msg := fmt.Sprintf("%d errors happened during parsing environment variables:\n", len(e.errors))
	for _, err := range e.errors {
		msg += fmt.Sprintf("\t- %s\n", err.Error())
	}
	return msg
}

func (e *ErrorList) IsEmpty() bool {
	if e.errors == nil || len(e.errors) == 0 {
		return true
	}

	return false
}

func (e *ErrorList) Append(err error) {
	e.errors = append(e.errors, err)
}

func (e *ErrorList) Join(errorList *ErrorList) {
	for _, err := range errorList.errors {
		e.errors = append(e.errors, err)
	}
}
