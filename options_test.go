package envparse

import "testing"

func TestSetPrefix(t *testing.T) {
	originalPrefix := DefaultPrefix

	testData := map[string]string{
		"ABC":     "ABC",
		"def":     "DEF",
		"_smth1_": "SMTH1",
		"3app":    "3APP",
	}

	for input, expectedResult := range testData {
		SetPrefix(input)
		if DefaultPrefix != expectedResult {
			t.Errorf("input '%s', expected '%s', but got '%s'", input, expectedResult, DefaultPrefix)
		}
	}

	DefaultPrefix = originalPrefix
}

func TestSetMaxDepth(t *testing.T) {
	originalMaxDepth := DefaultMaxDepth

	testData := map[int]int{
		1: 1,
		2: 2,
	}

	for input, expectedResult := range testData {
		SetMaxDepth(input)
		if DefaultMaxDepth != expectedResult {
			t.Errorf("input '%d', expected '%d', but got '%d'", input, expectedResult, DefaultMaxDepth)
		}
	}

	DefaultMaxDepth = originalMaxDepth
}
