package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func createDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func stringToByte(d deck) []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, stringToByte(d), 0666)
}
