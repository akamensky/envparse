package envparse

import (
	"reflect"
	"testing"
)

func TestNewEnvMap(t *testing.T) {
	type testDataType struct {
		in  []string
		out envMap
	}
	testData := []*testDataType{
		{
			in:  nil,
			out: map[string]string{},
		},
		{
			in: []string{
				"TEST=1",
			},
			out: map[string]string{
				"TEST": "1",
			},
		},
		{
			in: []string{
				"TEST/1=1",
			},
			out: map[string]string{
				"TEST/1": "1",
			},
		},
		{
			in: []string{
				"TEST",
			},
			out: map[string]string{},
		},
	}

	for _, datum := range testData {
		out := newEnvMap("", datum.in)
		if !reflect.DeepEqual(out, datum.out) {
			t.Errorf("expected '%v', but got '%v'", datum.out, out)
		}
	}
}

func TestEnvMap_Exists(t *testing.T) {
	m := envMap{
		"TEST": "1",
	}
	if !m.Exists("TEST") {
		t.Errorf("expected TEST exists, but got TEST not exists")
	}
	if m.Exists("TEST2") {
		t.Errorf("expected TEST2 not exists, but got TEST2 exists")
	}
	if m.Exists("TE") {
		t.Errorf("expected TE not exists, but got TE exists")
	}
}

func TestEnvMap_PrefixExists(t *testing.T) {
	m := envMap{
		"TEST": "1",
	}
	if !m.PrefixExists("TEST") {
		t.Errorf("expected prefix TEST exists, but got prefix TEST not exists")
	}
	if !m.PrefixExists("TE") {
		t.Errorf("expected prefix TE exists, but got prefix TE not exists")
	}
	if m.PrefixExists("TEST2") {
		t.Errorf("expected prefix TEST2 not exists, but got prefix TEST2 exists")
	}
}

func TestEnvMap_Get(t *testing.T) {
	m := envMap{
		"TEST": "1",
	}
	if m.Get("TEST") != "1" {
		t.Errorf("expected '%s', but got '%s'", "1", m.Get("TEST"))
	}
	if m.Get("TE") != "" {
		t.Errorf("expected '%s', but got '%s'", "", m.Get("TE"))
	}
	if m.Get("TEST2") != "" {
		t.Errorf("expected '%s', but got '%s'", "", m.Get("TEST2"))
	}
}

func TestEnvMap_GetPrefix(t *testing.T) {
	m := envMap{
		"TEST_":  "1",
		"TEST_1": "1",
		"TE_ST1": "1",
	}
	type testDataType struct {
		prefix string
		out    envMap
	}
	testData := []*testDataType{
		{
			prefix: "TEST",
			out: map[string]string{
				"":  "1",
				"1": "1",
			},
		},
		{
			prefix: "TE",
			out: map[string]string{
				"ST1": "1",
			},
		},
		{
			prefix: "TES_T2",
			out:    map[string]string{},
		},
	}
	for _, datum := range testData {
		out := m.GetPrefix(datum.prefix)
		if !reflect.DeepEqual(datum.out, out) {
			t.Errorf("expected '%v', but got '%v'", datum.out, out)
		}
	}
}

func TestEnvMap_GetSlicePrefixes(t *testing.T) {
	type testDataType struct {
		in  envMap
		out []string
	}
	testData := []*testDataType{
		{
			in: map[string]string{
				"HELLO": "1",
				"00":    "123",
				"100":   "abc",
				"_11":   "abc",
			},
			out: []string{"00", "100"},
		},
	}
	for _, datum := range testData {
		out := datum.in.GetSlicePrefixes()
		if !reflect.DeepEqual(datum.out, out) {
			t.Errorf("expected '%v', but got '%v'", datum.out, out)
		}
	}
}
