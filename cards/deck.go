package main

import (
	"fmt"
	"io/ioutil"
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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error while reading file", err)
	}

	dk := strings.Split(string(bs), ",")
	return dk

}
