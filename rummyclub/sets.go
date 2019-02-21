package rummyclub

type setCounter struct {
	rank    [27]int16
	suits   [27][4]bool
	indexes [27][4]int
	found   [27]int
}

func FindSets(hand []Card) []int {
	res := make([]int, 0)
	setsFound := setCounter{}

	for cardIndex, card := range hand {
		for index := 0; index < 27; index++ {
			if setsFound.rank[index] == 0 {
				setsFound.rank[index] = card.Rank
				setsFound.suits[index][card.Suit] = true
				setsFound.indexes[index][0] = cardIndex
				break
			} else {
				if setsFound.rank[index] == card.Rank {
					setsFound.suits[index][card.Suit] = true
					setsFound.found[index]++
					setsFound.indexes[index][setsFound.found[index]] = cardIndex
					break
				}
			}
		}
	}

	for index := 0; index < 27; index++ {
		if setsFound.rank[index] != 0 {
			count := 0
			for boolIndex := 0; boolIndex < 4; boolIndex++ {
				if setsFound.suits[index][boolIndex] {
					count++
				}
			}

			if count > 2 {
				for boolIndex := 0; boolIndex < 4; boolIndex++ {
					if setsFound.suits[index][boolIndex] {
						res = append(res, setsFound.indexes[index][boolIndex])
					}
				}
			}
		}
	}

	return res
}
