package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" // First init var using :=
	//card = "Five of Diamonds"
	card := newCard()

	fmt.Println(card)
	printState()
}

func newCard() string {
	return "Five of Diamonds"
}
