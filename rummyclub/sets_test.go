package rummyclub

import (
	"testing"
)

func TestFindNoSets(t *testing.T) {
	deck := MakeDeck()
	hand := deck[0:4]
	sets := FindSets(hand)
	if len(sets) != 0 {
		t.Errorf("No-sets cards found %d sets.\n", len(sets))
	}
}

func TestFindThreeCardSet(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 4)
	hand[0] = deck[0]
	hand[1] = deck[26]
	hand[2] = deck[52]
	hand[3] = deck[53]
	sets := FindSets(hand)
	if len(sets) != 3 {
		t.Errorf("One-sets cards (%v) found %d sets.\n", hand, len(sets))
	}
	correctIndexes := sets[0] == 0 && sets[1] == 1 && sets[2] == 2
	if !correctIndexes {
		t.Errorf("One-sets cards (%v) found wrong indexes (%v).\n", hand, sets)
	}
}
