package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inputFile = "input"
)

type Card struct {
	Id               int
	WinningNumbers   []int
	ScratchedNumbers []int
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
