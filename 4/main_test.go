package main

import "testing"

type CardTest struct {
	Input    string
	Expected Card
}

func (ct CardTest) ParseCard() (Card, bool) {
	card, err := NewCardFromString(ct.Input)
	if err != nil {
		panic(err)
	}
	if card.Id != ct.Expected.Id {
		return card, false
	}
	return card, true
}

type WinningTest struct {
	Input           string
	ExpectedNumbers []int
	ExpectedValue   int
}

func (wt WinningTest) CheckWinner(card Card) bool {
	if wt.ExpectedValue != card.Value() {
		return false
	}
	return true
}

func TestCardParser(t *testing.T) {
	tests := []CardTest{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Card{1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Card{2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Card{3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}}},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Card{4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}}},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Card{5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}}},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Card{6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}}},
	}

	for _, test := range tests {
		if card, ok := test.ParseCard(); !ok {
			t.Errorf("got: %+v - wanted %+v", card, test.Expected)
		}
	}
}

func TestWinningNumbers(t *testing.T) {
	tests := []WinningTest{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", []int{48, 83, 17, 86}, 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", []int{32, 61}, 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", []int{1, 21}, 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", []int{84}, 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", []int{}, 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", []int{}, 0},
	}

	for _, test := range tests {
		card, err := NewCardFromString(test.Input)
		if err != nil {
			panic(err)
		}
		if !test.CheckWinner(card) {
			t.Errorf("got: %+v, wanted: %+v", card.Value(), test.ExpectedValue)
		}
		if card.Value() != test.ExpectedValue {
			t.Errorf("card value - got: %d, wanted: %d", card.Value(), test.ExpectedValue)
		}
	}
}
