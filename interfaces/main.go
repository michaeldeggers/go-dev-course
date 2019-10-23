package main

import "fmt"

// interfaces used to define a function set
type bot interface { // interface type - can't create values out of this
	getGreeting() string
}

// these automatically become members of bot interface
type englishBot struct{} // concrete type
type spanishBot struct{} // concrete type

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
