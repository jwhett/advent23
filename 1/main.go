package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

var (
	writtenNumbers = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
)

type Number struct {
	Value, Index int
}

// parseNumbersFromString will return a list of integers found
// in input as a slice. Error is reported when no integers are
// found in input.
func parseNumbersFromString(input string) ([]Number, error) {
	found := make([]Number, 0)
	// Parse for regular ints
	for i, c := range input {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		found = append(found, Number{num, i})
	}
	// Parse for written numbers
	for _, wr := range writtenNumbers {
		if index := strings.Index(input, wr); index >= 0 {
			value, err := strconv.Atoi(wr)
			if err != nil {
				return found, err
			}
			found = append(found, Number{value, index})
		}
	}
	if len(found) == 0 {
		return found, fmt.Errorf("failed to parse numbers from input")
	}
	return found, nil
}

// parseCalibrationNumbers will return only the "calibration number"
// from a given line of input.
func parseCalibrationNumbers(input string) int {
	nums, err := parseNumbersFromString(input)
	if err != nil {
		return 0
	}
	// Note: Lines with a single digit will have that digit
	// treated as the first and the last digit.
	combined, err := strconv.Atoi(fmt.Sprintf("%d%d", nums[0].Value, nums[len(nums)-1].Value))
	if err != nil {
		return 0
	}
	return combined
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += parseCalibrationNumbers(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	fmt.Printf("Challenge 1 - Calibration total: %d\n", sum)
}
