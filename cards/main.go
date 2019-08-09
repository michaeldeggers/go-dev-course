package main

import "fmt"

// understood that this both creates and reads the same file - for learning
func main() {
	cards := newDeck()
	cards.shuffleCards()
	err := cards.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
	newCards := newDeckFromFile("my_cards")
	newCards.print()
}
