package rummyclub

import (
	"testing"

	"github.com/go-test/deep"
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
	hand[0] = deck[Y_1_1]
	hand[1] = deck[B_1_1]
	hand[2] = deck[R_1_1]
	hand[3] = deck[K_2_1]
	sets := FindSets(hand)
	if len(sets) != 3 {
		t.Errorf("One-sets cards (%v) found %d sets.\n", hand, len(sets))
	}
	correctIndexes := []int{0, 1, 2}
	if diff := deep.Equal(sets, correctIndexes); diff != nil {
		t.Errorf("One-sets cards (%v) found wrong indexes (%v).\n", hand, sets)
	}
}

func TestFindThreeCardSetMustBeDifferentSuits(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 4)
	hand[0] = deck[Y_1_1]
	hand[1] = deck[B_1_1]
	hand[2] = deck[Y_1_2]
	hand[3] = deck[K_2_1]
	sets := FindSets(hand)
	if len(sets) != 0 {
		t.Errorf("One-sets cards (%v) found %d sets.\n", hand, len(sets))
	}
}

func TestFindSevenCardSet(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 7)
	hand[0] = deck[Y_1_1]
	hand[1] = deck[R_2_1]
	hand[2] = deck[R_1_1]
	hand[3] = deck[K_2_1]
	hand[4] = deck[B_1_1]
	hand[5] = deck[B_2_1]
	hand[6] = deck[Y_2_1]
	sets := FindSets(hand)
	if len(sets) != 7 {
		t.Errorf("Two-sets cards (%v) found %d sets.\n", hand, len(sets))
	}
	correctIndexes := []int{0, 2, 4, 1, 3, 5, 6}
	if diff := deep.Equal(sets, correctIndexes); diff != nil {
		t.Errorf("One-sets cards (%v) found wrong indexes (%v).\n", hand, sets)
	}
}
