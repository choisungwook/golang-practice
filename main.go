package main

import "fmt"

func main() {
	cards := newDeck()
	// fmt.Println(cards.toString())
	// hand, remain := deal(cards, 5)
	// hand.print()
	// remain.print()
	cards.saveToFile("result.txt")

	new_deck := newDeckFromFile("result.txt")
	new_deck.print()

	fmt.Println("")
	new_deck.shuffle()
	new_deck.print()
}
