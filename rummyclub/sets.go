package rummyclub

type setCounter struct {
	rank    [27]int16
	suits   [27][SUIT_COUNT]bool
	indexes [27][4]int
	found   [27]int
}

func FindSets(hand []Card) []int {
	res := make([]int, 0)
	allSetsFound := setCounter{}

	// find
	for cardIndex, card := range hand {
		for index := 0; index < 27; index++ {
			if allSetsFound.rank[index] == 0 {
				allSetsFound.rank[index] = card.Rank
				allSetsFound.suits[index][card.Suit] = true
				allSetsFound.indexes[index][0] = cardIndex
				break
			} else {
				if allSetsFound.rank[index] == card.Rank {
					if !allSetsFound.suits[index][card.Suit] {
						allSetsFound.suits[index][card.Suit] = true
						allSetsFound.found[index]++
						allSetsFound.indexes[index][allSetsFound.found[index]] = cardIndex
					}
					break
				}
			}
		}
	}

	// copy actual sets
	for index := 0; index < 27; index++ {
		if allSetsFound.rank[index] == 0 {
			break
		}

		count := 0
		for boolIndex := 0; boolIndex < 4; boolIndex++ {
			if allSetsFound.suits[index][boolIndex] {
				count++
			}
		}

		if count > 2 {
			for suitsIndex := 0; suitsIndex < count; suitsIndex++ {
				res = append(res, allSetsFound.indexes[index][suitsIndex])
			}
		}
	}

	return res
}
