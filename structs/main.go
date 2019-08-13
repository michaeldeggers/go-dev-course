package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // we dont' have to specify the field name for embedded structs
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	frank := person{
		firstName: "Frank",
		lastName:  "Bro",
		contactInfo: contactInfo{
			email:   "frank@email.com",
			zipCode: 12345,
		},
	}

	frank.updateName("Jack")
	frank.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
