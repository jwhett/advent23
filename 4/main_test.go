package main

import "testing"

type CardTest struct {
	Input    string
	Expected Card
}

func (ct CardTest) ParseCard() bool {
	card, err := NewCardFromString(ct.Input)
	if err != nil {
		panic(err)
	}
	if card.Id != ct.Expected.Id {
		return false
	}
	return true
}

func TestCardParser(t *testing.T) {
	tests := []CardTest{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Card{}},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Card{}},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Card{}},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Card{}},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Card{}},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Card{}},
	}

	for _, test := range tests {
		if !test.ParseCard() {
			t.Errorf("input: %s - wanted %+v", test.Input, test.Expected)
		}
	}
}
