package envparse

import "fmt"

// ErrorList implements `error` interface with ability to combine multiple errors togethers
type ErrorList struct {
	errors []error
}

// FullError will return a new error made using `(*ErrorList).FullError()` or will return original `error` otherwise
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

// Error is needed to implement `error` interface
func (e *ErrorList) Error() string {
	return fmt.Sprintf("%d errors happened during parsing environment variables", len(e.errors))
}

// FullError combined all errors together in single message
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

// IsEmpty reports if *ErrorList is empty (0 appended errors)
func (e *ErrorList) IsEmpty() bool {
	if e.errors == nil || len(e.errors) == 0 {
		return true
	}

	return false
}

// Append adds a new error to the stack
func (e *ErrorList) Append(err error) {
	e.errors = append(e.errors, err)
}
