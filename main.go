package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/csterritt/rummyclub/rummyclub"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Cards:")
	deck := rummyclub.MakeDeck()
	rummyclub.Shuffle(deck)
	hand := deck[0:14]
	rummyclub.HandSort(hand)
	fmt.Printf("%d cards in the deck, %d in the hand:\n", len(deck), len(hand))
	for _, card := range hand {
		fmt.Printf("%s   ", card)
	}
	fmt.Println()
}
