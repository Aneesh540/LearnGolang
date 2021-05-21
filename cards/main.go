package main

import "fmt"

func main() {

	cards := newDeckFromFile("aneesh.txt")
	fmt.Println(len(cards))
	// cards := newDeck()
	fmt.Println(cards.toString())
	cards.print()
	// cards.saveToFile("aneesh.txt")

}
