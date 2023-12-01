package main

import (
	"slices"
	"testing"
)

type ParseNumbersTest struct {
	Input    string
	Expected []Number
}

func (pnt ParseNumbersTest) Test() (found []Number, expected bool) {
	found, err := parseNumbersFromString(pnt.Input)
	if err != nil {
		return found, false
	}
	match := slices.CompareFunc(found, pnt.Expected, func(x, y Number) int {
		if x.Value < y.Value {
			return -1
		} else if x.Value > y.Value {
			return 1
		}
		return 0
	})
	if match == 0 {
		expected = true
	} else {
		expected = false
	}
	return found, expected
}

type ParseCalibrationTest struct {
	Input    string
	Expected int
}

func (pct ParseCalibrationTest) Test() (found int, expected bool) {
	found = parseCalibrationNumbers(pct.Input)
	if found == pct.Expected {
		expected = true
	} else {
		expected = false
	}
	return found, expected
}

func TestParseNumbersFromString(t *testing.T) {
	tests := []ParseNumbersTest{
		{"1abc2", []Number{{1, 0}, {2, 4}}},
		{"pqr3stu8vwx", []Number{{3, 3}, {8, 7}}},
		{"a1b2c3d4e5f", []Number{{1, 1}, {2, 3}, {3, 5}, {4, 7}, {5, 9}}},
		{"treb7uchet", []Number{{7, 4}}},
	}

	for _, test := range tests {
		if got, result := test.Test(); !result {
			t.Errorf("input: %+v. got: %d, wanted %d", test.Input, got, test.Expected)
		}
	}
}

func TestParseCalibrationNumbers(t *testing.T) {
	tests := []ParseCalibrationTest{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		if got, result := test.Test(); !result {
			t.Errorf("input: %+v. got: %d, wanted %d", test.Input, got, test.Expected)
		}
	}
}
