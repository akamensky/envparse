package envparse

import "fmt"

type errorList struct {
	errors []error
}

func newErrorList() *errorList {
	return &errorList{errors: make([]error, 0)}
}

func (e *errorList) Error() string {
	return fmt.Sprintf("%d errors happened during parsing environment variables", len(e.errors))
}

func (e *errorList) FullError() string {
	msg := e.Error()
	if len(e.errors) > 0 {
		msg = fmt.Sprintf("%s:", msg)
	}
	for _, err := range e.errors {
		msg += fmt.Sprintf("\n\t- %s", err.Error())
	}
	return msg
}

func (e *errorList) IsEmpty() bool {
	if e.errors == nil || len(e.errors) == 0 {
		return true
	}

	return false
}

func (e *errorList) Append(err error) {
	e.errors = append(e.errors, err)
}
