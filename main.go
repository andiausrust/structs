package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	andi := person{
		firstName: "andi",
		lastName: "mayer",
		contact: contactInfo{
			email: "amayer@test.com",
			zipCode: 12323},
	}

	fmt.Println(andi)

	andi.firstName = "Andi"

	andi.print()
}

func (p person) print(){
	fmt.Printf("%+v",p)
}

func (p person) updateName (nfn string) {
	p.firstName = nfn
}