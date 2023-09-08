package main

func main() {

	cards := deck{"Ace of Diamond", newCards()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCards() string {
	return "Five of Diamonds"
}
