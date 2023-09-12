package main

func main() {
	//cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// cards.print()
	// hand.print()
	// remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}