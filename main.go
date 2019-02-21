package main

import (
	"fmt"

	"github.com/csterritt/rummyclub/rummyclub"
)

func main() {
	fmt.Println("Cards:")
	deck := rummyclub.MakeDeck()
	for _, card := range deck {
		fmt.Printf("%s   ", card)
	}
	fmt.Println()
}
