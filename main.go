package main

import (
	"fmt"
	"github.com/csterritt/rummyclub/rummyclub"
)

func main() {
	fmt.Println("Cards:")
	card := rummyclub.Card{Suit: rummyclub.YELLOW, Rank: 1}
	fmt.Println("1 yellow:", card)
	card = rummyclub.Card{Suit: rummyclub.RED, Rank: 9}
	fmt.Println("9 red:", card)
	card = rummyclub.Card{Suit: rummyclub.BLUE, Rank: 10}
	fmt.Println("10 blue:", card)
	card = rummyclub.Card{Suit: rummyclub.BLACK, Rank: 13}
	fmt.Println("13 black:", card)
	card = rummyclub.Card{Suit: rummyclub.WILD, Rank: 13}
	fmt.Println("Wildcard:", card)
}
