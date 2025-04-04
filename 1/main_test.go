package main

import (
	"slices"
	"testing"
)

type ParseNumbersTest struct {
	Input    string
	Expected []int
}

func (pnt ParseNumbersTest) Test() (found []int, expected bool) {
	found, err := parseNumbersFromString(pnt.Input)
	if err != nil {
		return found, false
	}
	if match := slices.Compare(found, pnt.Expected); match == 0 {
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
		{"1abc2", []int{1, 2}},
		{"pqr3stu8vwx", []int{3, 8}},
		{"a1b2c3d4e5f", []int{1, 2, 3, 4, 5}},
		{"treb7uchet", []int{7}},
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
