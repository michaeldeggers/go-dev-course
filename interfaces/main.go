package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printEnglishGreeting(eb)
	printSpanishGreeting(sb)
}

func printEnglishGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printSpanishGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
