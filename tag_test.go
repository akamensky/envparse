package envparse

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	var in string
	var expectedOut *tagType
	var expectedErr error

	in = "name=HELLO1234_*\n/"
	expectedOut = &tagType{name: "HELLO1234", required: false, defaultValue: ""}
	expectedErr = nil
	tag, err := parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,required"
	expectedOut = &tagType{name: "HELLO1234", required: true, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default"
	expectedOut = &tagType{name: "HELLO1234", required: false, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default=12"
	expectedOut = &tagType{name: "HELLO1234", required: false, defaultValue: "12"}
	expectedErr = nil
	tag, err = parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default=,required"
	expectedOut = &tagType{name: "HELLO1234", required: true, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=HELLO1234_*\n/,default=12,required"
	expectedOut = nil
	expectedErr = tagErrIncompatibleFields
	tag, err = parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}
}

func TestParseTag2(t *testing.T) {
	// Test allowing underscores in `name`
	var in string
	var expectedOut *tagType
	var expectedErr error

	in = "name=HELLO_1234"
	expectedOut = &tagType{name: "HELLO_1234", required: false, defaultValue: ""}
	expectedErr = nil
	tag, err := parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}

	in = "name=_HELLO_1234_"
	expectedOut = &tagType{name: "HELLO_1234", required: false, defaultValue: ""}
	expectedErr = nil
	tag, err = parseTag(in, "")
	if !reflect.DeepEqual(tag, expectedOut) {
		t.Errorf("expected tag to be %v, but got %v", expectedOut, tag)
	}
	if expectedErr != err {
		t.Errorf("expected error to be '%s', but got '%s'", expectedErr, err)
	}
}
