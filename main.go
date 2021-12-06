package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	hand, remain := deal(cards, 5)
	hand.print()
	remain.print()
}
