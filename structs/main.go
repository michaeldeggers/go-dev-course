package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// first way -- danger in not specifycing first / last fields
	jack := person{"Jack", "Bro", contactInfo{"", 0}}
	fmt.Printf("%+v", jack)

	// second way -- better -- can reorder fields later without losing proper def
	mike := person{firstName: "Mike", lastName: "Bro"}
	fmt.Printf("%+v", mike)

	// third way - initialize first, define later
	var bob person
	bob.firstName = "Bob"
	bob.lastName = "Bro"
	fmt.Printf("%+v", bob)

	frank := person{
		firstName: "Frank",
		lastName:  "Bro",
		contact: contactInfo{
			email:   "frank@email.com",
			zipCode: 12345,
		},
	}
	fmt.Println(frank)
}
