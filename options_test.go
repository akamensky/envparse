package envparse

import "testing"

func TestSetPrefix(t *testing.T) {
	originalPrefix := defaultPrefix

	testData := map[string]string{
		"ABC":     "ABC",
		"def":     "DEF",
		"_smth1_": "SMTH1",
		"3app":    "3APP",
	}

	for input, expectedResult := range testData {
		SetPrefix(input)
		if defaultPrefix != expectedResult {
			t.Errorf("input '%s', expected '%s', but got '%s'", input, expectedResult, defaultPrefix)
		}
	}

	defaultPrefix = originalPrefix
}

func TestSetMaxDepth(t *testing.T) {
	originalMaxDepth := defaultMaxDepth

	testData := map[int]int{
		1: 1,
		2: 2,
	}

	for input, expectedResult := range testData {
		SetMaxDepth(input)
		if defaultMaxDepth != expectedResult {
			t.Errorf("input '%d', expected '%d', but got '%d'", input, expectedResult, defaultMaxDepth)
		}
	}

	defaultMaxDepth = originalMaxDepth
}

func TestSetUnsetEnv(t *testing.T) {
	originalUnsetEnv := unsetEnv

	SetUnsetEnv(true)
	if !unsetEnv {
		t.Errorf("expected unsetEnv == true, but got unsetEnv == false")
	}

	unsetEnv = originalUnsetEnv
}
