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

func TestFindSetWithWildcard(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 14)
	hand[0] = deck[Y_3_1]
	hand[1] = deck[Y_4_1]
	hand[2] = deck[Y_5_1]
	hand[3] = deck[Y10_1]
	hand[4] = deck[B_2_1]
	hand[5] = deck[B_4_1]
	hand[6] = deck[B12_1]
	hand[7] = deck[B12_2]
	hand[8] = deck[R_2_1]
	hand[9] = deck[R_4_1]
	hand[10] = deck[R10_1]
	hand[11] = deck[R12_1]
	hand[12] = deck[K_8_1]
	hand[13] = Card{Suit: WILD, Rank: 0}
	sets := FindSets(hand)
	if len(sets) != 3 {
		t.Errorf("Set-with-wildcard cards (%v) found %d sets.\n", hand, len(sets))
	}
	correctIndexes := []int{1, 5, 9}
	if diff := deep.Equal(sets, correctIndexes); diff != nil {
		t.Errorf("Set-with-wildcard cards (%v) found wrong indexes (%v).\n", hand, sets)
	}
}

func TestFindThreeCardSetMustBeDifferentIndexes(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 14)
	hand[0] = deck[Y_2_1]
	hand[1] = deck[Y_7_1]
	hand[2] = deck[Y11_1]
	hand[3] = deck[B_5_1]
	hand[4] = deck[B_5_2]
	hand[5] = deck[B13_1]
	hand[6] = deck[R_2_1]
	hand[7] = deck[R_4_1]
	hand[8] = deck[R_7_1]
	hand[9] = deck[R10_1]
	hand[10] = deck[R11_1]
	hand[11] = deck[K_1_1]
	hand[12] = deck[K_2_1]
	hand[13] = deck[K10_1]
	sets := FindSets(hand)
	if len(sets) != 3 {
		t.Errorf("Set-with-different-indexes cards (%v) found %d sets.\n", hand, len(sets))
	}
	correctIndexes := []int{0, 6, 12}
	if diff := deep.Equal(sets, correctIndexes); diff != nil {
		t.Errorf("Set-with-different-indexes cards (%v) found wrong indexes (%v).\n", hand, sets)
	}
}

func TestFindThreeCardSetWithRepeatedCard(t *testing.T) {
	deck := MakeDeck()
	hand := make([]Card, 14)
	hand[0] = deck[Y_2_1]
	hand[1] = deck[Y_7_1]
	hand[2] = deck[Y11_1]
	hand[3] = deck[B_5_1]
	hand[4] = deck[B_5_2]
	hand[5] = deck[R_2_1]
	hand[6] = deck[R_2_2]
	hand[7] = deck[R_4_1]
	hand[8] = deck[R_7_1]
	hand[9] = deck[R10_1]
	hand[10] = deck[R11_1]
	hand[11] = deck[K_1_1]
	hand[12] = deck[K_2_1]
	hand[13] = deck[K10_1]
	sets := FindSets(hand)
	if len(sets) != 6 {
		t.Errorf("Set-with-repeated-card cards (%v) found %d sets.\n", hand, len(sets))
	}
	correctIndexes := []int{0, 5, 12, 0, 6, 12}
	if diff := deep.Equal(sets, correctIndexes); diff != nil {
		t.Errorf("Set-with-repeated-card cards (%v) found wrong indexes (%v).\n", hand, sets)
	}
}
