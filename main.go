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
	for _, card := range deck {
		fmt.Printf("%s   ", card)
	}
	fmt.Println()
}
