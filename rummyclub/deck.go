package rummyclub

import "math/rand"

func MakeDeck() []Card {
	res := make([]Card, 106)
	index := 0
	for suit := YELLOW; suit < WILD; suit++ {
		for ignore := 0; ignore < 2; ignore++ {
			for rank := 1; rank < 14; rank++ {
				res[index].Suit = suit
				res[index].Rank = int16(rank)
				index++
			}
		}
	}
	res[index].Suit = WILD
	index++
	res[index].Suit = WILD
	index++

	return res
}

func Shuffle(slc []Card) {
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
}
