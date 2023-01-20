package envparse

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	var in string
	var expectedOut *tagType
	expectedErr := tagErrNameField
	tag, err := parseTag(in)
	if tag != expectedOut {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name="
	expectedOut = nil
	expectedErr = tagErrNameField
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/"
	expectedOut = &tagType{name: "HELLO1234", required: false, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,required"
	expectedOut = &tagType{name: "HELLO1234", required: true, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default"
	expectedOut = &tagType{name: "HELLO1234", required: false, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default=12"
	expectedOut = &tagType{name: "HELLO1234", required: false, defaultValue: "12"}
	expectedErr = nil
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default=,required"
	expectedOut = &tagType{name: "HELLO1234", required: true, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default=12,required"
	expectedOut = nil
	expectedErr = tagErrIncompatibleFields
	tag, err = parseTag(in)
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}
}
