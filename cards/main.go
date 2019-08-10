package main

import "fmt"

// understood that this both creates and reads the same file - for learning
func main() {
	filename := "demo_files/my_cards"
	cards := newDeck()
	cards.shuffleCards()
	err := cards.saveToFile(filename)
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
	newCards := newDeckFromFile(filename)
	newCards.print()
}
