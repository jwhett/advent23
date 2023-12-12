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

type Card struct {
	Id               int
	WinningNumbers   []int
	ScratchedNumbers []int
}

func (c Card) MatchingNumbers() []int {
	return []int{}
}

func (c Card) IsAWinner() bool {
	return false
}

// Value returns the point value of a card given
// its matching numbers.
func (c Card) Value() int {
	return 0
}

// atoiSlice applies strconv.Atoi to each element
// of its input and returns a slice of their int
// values.
func atoiSlice(writtenNumber []string) []int {
	var convertedNumbers []int
	for _, strnum := range writtenNumber {
		num, err := strconv.Atoi(strnum)
		if err != nil {
			panic(err)
		}
		convertedNumbers = append(convertedNumbers, num)
	}
	return convertedNumbers
}

func parseCardId(in string) (int, error) {
	fields := strings.Fields(in)
	rawId, err := strconv.Atoi(strings.ReplaceAll(fields[1], ":", ""))
	if err != nil {
		return -1, err
	}
	return rawId, nil
}

func parseWinningNumbers(in string) []int {
	rawNumberString := strings.Split(strings.Split(in, " | ")[0], ": ")
	strList := strings.Fields(rawNumberString[len(rawNumberString)-1])
	return atoiSlice(strList)
}

func parseScratchedNumbers(in string) []int {
	fields := strings.Split(in, " | ")
	scratchedNumbers := strings.Fields(fields[len(fields)-1])
	return atoiSlice(scratchedNumbers)
}

func NewCardFromString(in string) (Card, error) {
	id, err := parseCardId(in)
	if err != nil {
		return Card{}, err
	}
	return Card{id, parseWinningNumbers(in), parseScratchedNumbers(in)}, nil
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
