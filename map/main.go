package main

import (
	"fmt"
)

func main() {
	// map[<key type>]<value type>
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
