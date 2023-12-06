package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "input"
)

// Positions of each color.
const (
	REDPOS int = iota
	GREENPOS
	BLUEPOS
)

// Bag is a container from which we pull Rounds from. The
// fields represent the maximum number of each color cube.
type Bag struct {
	Red, Green, Blue int
}

func (b Bag) IsPOssibleGame(game Game) bool {
	for _, round := range game.Rounds {
		if round.Red > b.Red || round.Green > b.Green || round.Blue > b.Blue {
			return false
		}
	}
	return true
}

// Rounds are handfuls of cubes. The three fields represent
// how many of each color were drawn from that round.
//
// Example: 3 blue, 4 red
// In this example, no green cubes were pulled from the Bag.
type Round struct {
	Red, Green, Blue int
}

// Game represents a series of rounds. Each Game has its own
// ID and some number of rounds that were played.
//
// Example: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// In this example, three Rounds were played in Game 1.
type Game struct {
	Id     int
	Rounds []Round
}

// FewestCubePower calculates the fewest cubes necessary
// to play a given game then returns the "power" which
// is defined as the product of the fewest Red, Green,
// and Blue cubes.
func (g Game) FewestCubePower() int {
	neededRed, neededGreen, neededBlue := -1, -1, -1
	for _, round := range g.Rounds {
		if neededRed == -1 || round.Red > neededRed {
			neededRed = round.Red
		}
		if neededGreen == -1 || round.Green > neededGreen {
			neededGreen = round.Green
		}
		if neededBlue == -1 || round.Blue > neededBlue {
			neededBlue = round.Blue
		}
	}
	return neededRed * neededGreen * neededBlue
}

func parseGame(line string) Game {
	gameInfo := strings.Split(line, ": ")
	var gameId int
	fmt.Sscanf(gameInfo[0], "Game %d", &gameId)
	rounds := parseRounds(gameInfo[1])
	return Game{gameId, rounds}
}

func parseRounds(line string) []Round {
	rounds := make([]Round, 0)
	rawRounds := strings.Split(line, ";")
	for _, rawRound := range rawRounds {
		cubes := parseCubes(rawRound)
		rounds = append(rounds, Round{cubes[REDPOS], cubes[GREENPOS], cubes[BLUEPOS]})
	}
	return rounds
}

func parseCubes(input string) []int {
	found := make(map[string]int, 0)
	rawCubes := strings.Split(input, ", ")
	for set := range rawCubes {
		var count int
		var color string
		fmt.Sscanf(rawCubes[set], "%d %s", &count, &color)
		found[color] = count
	}
	return []int{found["red"], found["green"], found["blue"]}
}

var (
	bag = Bag{
		12,
		13,
		14,
	}
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var sumPossibleGames, sumPowers int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		if bag.IsPOssibleGame(game) {
			sumPossibleGames += game.Id
		}
		sumPowers += game.FewestCubePower()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v", err)
	}

	fmt.Printf("Sum of possible game IDs: %d\n", sumPossibleGames)
	fmt.Printf("Sum of cube powers: %d\n", sumPowers)
}
