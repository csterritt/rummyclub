package rummyclub

import (
	"testing"
)

func TestFindNoRuns(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 4)
	hand[0] = deck[Y_1_1]
	hand[1] = deck[B_1_1]
	hand[2] = deck[R_1_1]
	hand[3] = deck[K_2_1]
	runs := FindRuns(hand)
	if len(runs) != 0 {
		t.Errorf("No-runs cards found %d runs.\n", len(runs))
	}
}

func TestFindNoRuns2(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 4)
	hand = append(hand, deck[Y_7_1:Y_8_1]...)
	hand = append(hand, deck[Y11_1:Y13_1]...)
	runs := FindRuns(hand)
	if len(runs) != 0 {
		t.Errorf("No-runs-2 cards found %d runs.\n", len(runs))
	}
}

func TestFindNoRuns3(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 4)
	hand = append(hand, deck[Y12_2:B_3_1]...)
	runs := FindRuns(hand)
	if len(runs) != 0 {
		t.Errorf("No-runs-3 cards found %d runs.\n", len(runs))
	}
}

func TestFindOneRun(t *testing.T) {
	deck := MakeDeck()
	hand := deck[Y_1_1:Y_4_1]
	runs := FindRuns(hand)
	if len(runs) != 3 {
		t.Errorf("One-runs cards found %d runs.\n", len(runs))
	}
}

func TestFindTwoRuns(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 0)
	hand = append(hand, deck[Y_1_1])
	hand = append(hand, deck[Y_7_1:Y10_1]...)
	hand = append(hand, deck[B_4_1:B_8_1]...)
	hand = append(hand, deck[K_1_1])
	runs := FindRuns(hand)
	if len(runs) != 7 {
		t.Errorf("Two-runs cards found %d runs (%v).\n", len(runs), runs)
	}
}
