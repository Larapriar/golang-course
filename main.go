package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")

	// // hand, remainingDeck := deal(cards, 5)
	// // hand.print()
	// // fmt.Println("#######")
	// // remainingDeck.print()

	// s := newDeckFromFile("my_cards")
	// s.print()

}
