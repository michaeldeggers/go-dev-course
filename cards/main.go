package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string { // 'string' required to designate return type
	return "Five of Diamonds"
}
