package main
import "fmt"

//create a new type of 'deck'
//whicl is slice of stings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
