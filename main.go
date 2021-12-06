package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remain := deal(cards, 5)
	hand.print()

	fmt.Println("===")

	remain.print()
}
