package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// new deck returns a deck with string types, rather than an empty deck of type string
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// this is called a receiver
// 'd' refers to the 'copy' of deck accessing the method ( the instance )
// 'deck' is the 'type' this function is applied to
// very similar to the idea 'this' from javascript
// 'copy' by convention is usually one or two letters referencing the type being used
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
