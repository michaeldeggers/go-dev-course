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
	frank := person{
		firstName: "Frank",
		lastName:  "Bro",
		contact: contactInfo{
			email:   "frank@email.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v", frank)
}
