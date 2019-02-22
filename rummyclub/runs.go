package rummyclub

// NOTE: hand MUST be sorted, suit increasing first then rank increasing within each suit.
func FindRuns(hand []Card) []int {
	res := make([]int, 0)

	runStart := 0
	index := 1
	prevCard := hand[0]
	currentSuit := prevCard.Suit
	for index < len(hand) {
		card := hand[index]
		if card.Suit != currentSuit || card.Rank != prevCard.Rank+1 {
			if index-runStart > 2 {
				for runStart < index {
					res = append(res, runStart)
					runStart++
				}
			}
			runStart = index
			currentSuit = card.Suit
		}
		prevCard = card
		index++
	}

	// run at end of the hand?
	if len(hand)-runStart > 2 {
		for runStart < len(hand) {
			res = append(res, runStart)
			runStart++
		}
	}

	return res
}
