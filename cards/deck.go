package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	newdeck := deck{}

	cardType := []string{"Ace", "Spade", "diamond", "Club"}
	cardValue := []string{"A", "1", "2", "3"}

	for _, typ := range cardType {
		for _, val := range cardValue {
			newdeck = append(newdeck, typ+"-"+val)

		}
	}

	return newdeck

}

func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}
