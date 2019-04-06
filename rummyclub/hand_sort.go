package rummyclub

import "sort"

func HandSort(hand []Card) {
	sort.Slice(hand, func(i, j int) bool {
		if hand[i].Suit < hand[j].Suit {
			return true
		}
		if hand[i].Suit > hand[j].Suit {
			return false
		}
		return hand[i].Rank <= hand[j].Rank
	})
}
