package rummyclub

import "fmt"

type SuitT int

const (
	YELLOW SuitT = iota
	BLUE
	RED
	BLACK
	WILD
)

type Card struct {
	Suit SuitT
	Rank int
}

func (card Card) String() string {
	if card.Suit == WILD {
		return " * "
	}

	var suitChar string
	switch card.Suit {
	case YELLOW:
		suitChar = "Y"
	case BLUE:
		suitChar = "B"
	case RED:
		suitChar = "R"
	case BLACK:
		suitChar = "K"
	case WILD:
		suitChar = "*"
	}

	padChar := ""
	if card.Rank < 10 {
		padChar = "_"
	}

	return fmt.Sprintf("%s%s%d", suitChar, padChar, card.Rank)
}
