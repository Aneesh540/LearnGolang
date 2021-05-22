package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) updateName(newName string) {
	p.firstName = newName
	fmt.Println("name updated")
}

func main() {

	// GOLANG is PASS BY VALUE for all data types

	a := person{
		firstName: "unknown",
		lastName:  "DARK netflix",
		contactInfo: contactInfo{
			email: "gmail.com",
			zip:   324009,
		},
	}

	fmt.Println(a.contactInfo)
	a.updateName("bholu")
	fmt.Println(a)
}
