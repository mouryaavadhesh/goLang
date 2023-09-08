package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamond", newCards()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCards() string {
	return "Five of Diamonds"
}
