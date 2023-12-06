// Template file only!
// Saving some time each day to ensure I
// actually write tests.
package main

import "testing"

type GameTest struct {
	Field    string // this will likely be a Struct to be tested.
	Expected bool   // expected test result.
}

// Test only exists to prevent errors below.
// These should be implemented in the non-test
// package.
func (st GameTest) Test() bool {
	return true
}

var gameInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestParseCubes(t *testing.T) {
	cubes := parseCubes("3 blue, 4 red")
	expected := []int{4, 0, 3}
	for i := range cubes {
		if cubes[i] != expected[i] {
			t.Errorf("parseCubes: Failed to parse the correct nubmer of cubes. Expected: %v, Got: %v", expected, cubes)
		}
	}
}

func TestParseRound(t *testing.T) {
	rounds := parseRounds("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	if len(rounds) != 3 {
		t.Errorf("parseRound: Failed to parse the correct number of rounds. Expected: 3, Got: %d", len(rounds))
	}
	round := rounds[0]
	expectedRound := Round{4, 0, 3}
	if round.Red != expectedRound.Red || round.Green != expectedRound.Green || round.Blue != expectedRound.Blue {
		t.Errorf("parseRound: Failed to parse the number of cubes in round. Expected: %v, Got: %v", expectedRound, round)
	}
}

func TestParseGame(t *testing.T) {
	game := parseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	expected := Game{
		1,
		[]Round{
			{4, 0, 3},
			{1, 2, 6},
			{0, 2, 0},
		},
	}
	if len(game.Rounds) != len(expected.Rounds) {
		t.Errorf("parseGame: Unexpected number of rounds parsed. Expected: %d, Got: %d", len(expected.Rounds), len(game.Rounds))
	}
}

func TestIsPossibleGame(t *testing.T) {
	passingGame := parseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	impossibleGame := parseGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 redGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	if !bag.IsPOssibleGame(passingGame) {
		t.Error("IsPossibleGame: game was expected to pass.")
	}
	if bag.IsPOssibleGame(impossibleGame) {
		t.Error("IsPossibleGame: game was expected to fail.")
	}
}
