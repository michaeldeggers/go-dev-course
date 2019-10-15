package main

import (
	"fmt"
)

func main() {
	//var colors map[string]string

	colors := make(map[string]string)

	// key must be appropriate type when setting and referencing
	colors["white"] = "#ffffff"
	colors["black"] = "#000"

	// map[<key type>]<value type>
	//colors := map[string]string{
	//	"red": "#ff0000",
	//	"green": "#4bf745",
	//}

	delete(colors, "white")

	fmt.Println(colors)
}
