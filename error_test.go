package envparse

import (
	"fmt"
	"testing"
)

func TestNewErrorList(t *testing.T) {
	e := newErrorList()
	if len(e.errors) != 0 {
		t.Errorf("expected list of errors to be 0, but got %d", len(e.errors))
	}
}

func TestErrorList_Error(t *testing.T) {
	e := newErrorList()
	expected := "0 errors happened during parsing environment variables"
	if e.Error() != expected {
		t.Errorf("unexpected error message; expected '%s', but got '%s'", expected, e.Error())
	}
}

func TestErrorList_FullError(t *testing.T) {
	e := newErrorList()
	expected := "0 errors happened during parsing environment variables"
	if e.FullError() != expected {
		t.Errorf("unexpected error message; expected '%s', but got '%s'", expected, e.FullError())
	}

	e.errors = append(e.errors, fmt.Errorf("some error"))
	expected = fmt.Sprintf("1 errors happened during parsing environment variables:\n\t- some error")
	if e.FullError() != expected {
		t.Errorf("unexpected error message; expected '%s', but got '%s'", expected, e.FullError())
	}
}

func TestErrorList_IsEmpty(t *testing.T) {
	e := newErrorList()
	if e.IsEmpty() != true {
		t.Errorf("expected IsEmpty == true, but got IsEmpty == false")
	}

	e.errors = append(e.errors, fmt.Errorf("some error"))
	if e.IsEmpty() != false {
		t.Errorf("expected IsEmpty == false, but got IsEmpty == true")
	}
}

func TestErrorList_Append(t *testing.T) {
	e := newErrorList()
	e.Append(fmt.Errorf("some error"))
	if e.IsEmpty() != false {
		t.Errorf("expected IsEmpty == false, but got IsEmpty == true")
	}
}

func TestFullError(t *testing.T) {
	errorList := &ErrorList{errors: []error{
		fmt.Errorf("hello"),
		fmt.Errorf("world"),
	}}
	if FullError(errorList).Error() != errorList.FullError() {
		t.Errorf("expected '%v', but got '%v'", errorList.FullError(), FullError(errorList).Error())
	}

	err := fmt.Errorf("hello world")
	if FullError(err) != err {
		t.Errorf("expected '%v', but got '%v'", err, FullError(err))
	}
}
