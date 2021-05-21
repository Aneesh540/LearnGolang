package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	aneesh := person{"Aneesh", "Jain",contactInfo{"gmail", 324009}}
	nitish := person{firstName: "Nitish", lastName: "Jain", contact: contactInfo{"gmail", 324009}}
	fmt.Println(aneesh)
	fmt.Println(nitish)

	var a2 person
	fmt.Println(a2)
}
