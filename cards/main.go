package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string { // 'string' required to designate return type
	return "Five of Diamonds"
}
