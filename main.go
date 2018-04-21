package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	tolga := person{
		firstName: "Tolga",
		lastName:  "Boztuna",
		contactInfo: contactInfo{
			email:   "tboztuna@yahoo.com",
			zipCode: 34500,
		},
	}

	tolga.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
