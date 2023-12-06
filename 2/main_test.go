// Template file only!
// Saving some time each day to ensure I
// actually write tests.
package main

import "testing"

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

func TestFewestCubePower(t *testing.T) {
	game := parseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
	expected := 4 * 2 * 6
	if pow := game.FewestCubePower(); pow != expected {
		t.Errorf("FewestCubePower: failed to calculate expected power. Got: %d, Expected: %d", pow, expected)
	}
}
