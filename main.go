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

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}


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
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	// jim.print()
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
