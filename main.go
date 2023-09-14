package main

import "fmt"


type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}
func main() {
	//cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// cards.print()
	// hand.print()
	// remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	// cards.shuffle()
	// cards.print()
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}