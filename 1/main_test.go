package main

import "testing"

type ParseCalibrationTest struct {
	Input    string
	Expected int
	Pass     bool
}

func (st ParseCalibrationTest) Test() (found int, expected bool) {
	found = parseCalibrationNumbers(st.Input)
	if found == st.Expected {
		expected = true
	} else {
		expected = false
	}
	return found, expected
}

func TestParseCalibrationNumbers(t *testing.T) {
	tests := []ParseCalibrationTest{
		{"1abc2", 12, true},
		{"pqr3stu8vwx", 38, true},
		{"a1b2c3d4e5f", 15, true},
		{"treb7uchet", 77, true},
	}

	for _, test := range tests {
		if got, result := test.Test(); result != test.Pass {
			t.Errorf("input: %+v. got: %d, wanted %d", test.Input, got, test.Expected)
		}
	}
}
