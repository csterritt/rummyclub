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

	runs := rummyclub.FindRuns(hand)
	if len(runs) > 0 {
		fmt.Print("  Runs: (")
		for _, index := range runs {
			fmt.Print(index, " ")
		}
		fmt.Print(") ")
		for _, index := range runs {
			fmt.Print(hand[index], " ")
		}
		fmt.Println()
	}

	sets := rummyclub.FindSets(hand)
	if len(sets) > 0 {
		fmt.Print("  Sets: (")
		for _, index := range sets {
			fmt.Print(index, " ")
		}
		fmt.Print(") ")
		for _, index := range sets {
			fmt.Print(hand[index], " ")
		}
		fmt.Println()
	}
}
