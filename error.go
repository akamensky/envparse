package envparse

import "fmt"

type ErrorList struct {
	errors []error
}

func FullError(err error) error {
	if errorList, ok := err.(*ErrorList); ok {
		return fmt.Errorf(errorList.FullError())
	} else {
		return err
	}
}

func newErrorList() *ErrorList {
	return &ErrorList{errors: make([]error, 0)}
}

func (e *ErrorList) Error() string {
	return fmt.Sprintf("%d errors happened during parsing environment variables", len(e.errors))
}

func (e *ErrorList) FullError() string {
	msg := e.Error()
	if len(e.errors) > 0 {
		msg = fmt.Sprintf("%s:", msg)
	}
	for _, err := range e.errors {
		msg += fmt.Sprintf("\n\t- %s", err.Error())
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
