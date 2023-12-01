package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	inputFile = "input"
)

// parseNumbersFromString will return a list of integers found
// in input as a slice. Error is reported when no integers are
// found in input.
func parseNumbersFromString(input string) ([]int, error) {
	found := make([]int, 0)
	for _, c := range input {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		found = append(found, i)
	}
	if len(found) == 0 {
		return found, fmt.Errorf("failed to parse numbers from input.")
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
	combined, err := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}
}
