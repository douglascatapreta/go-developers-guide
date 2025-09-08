package main

func main() {
	d := newDeck()
	d.shuffle()
	d.print()

	hand, remainingDeck := deal(d, 5)
	hand.print()
	remainingDeck.print()

	d.saveToFile("my_cards")

	loadedDeck := newDeckFromFile("my_cards")
	loadedDeck.print()
}
