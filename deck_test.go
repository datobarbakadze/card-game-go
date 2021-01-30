package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()
	if len(d) != 16 {
		t.Errorf("Expceted deck to be length of 20 and actual is %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of ace of spades")
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := createDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards")
	}

	os.Remove(filename)
}
