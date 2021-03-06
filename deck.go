package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	cmStrings := string(bs)                      // convert byte to string
	sliceString := strings.Split(cmStrings, ",") // convert string to slice
	return deck(sliceString)                     // convert slice to deck and return the deck
}

func (d deck) shuffle() {
	currentTime := time.Now()
	newSource := rand.NewSource(currentTime.UnixNano())
	r := rand.New(newSource)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
