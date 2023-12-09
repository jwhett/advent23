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

func parseCardId(in string) (int, error) {
	fields := strings.Fields(in)
	rawId, err := strconv.Atoi(strings.ReplaceAll(fields[1], ":", ""))
	if err != nil {
		return -1, err
	}
	return rawId, nil
}

func parseWinningNumbers(in string) []int {
	return []int{}
}

func parseScratchedNumbers(in string) []int {
	return []int{}
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
