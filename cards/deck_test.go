package main

import (
	"os"
	"testing"
)

func TestShuffle(t *testing.T) {
	cardsOriginal := newDeck()
	cards := newDeck()
	cards.shuffle()
	sameLocation := 0
	for i, value := range cardsOriginal {
		if cards[i] == value {
			sameLocation++
		}
	}
	if sameLocation == len(cardsOriginal) {
		t.Errorf("Cards do not shuffle")
	}
}

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", cards[0])
	}
	if cards[len(cards)-1] != "Five of Clubs" {
		t.Errorf("Expected first card to be Five of Clubs, but got %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	// remove the file
	os.Remove("_decktesting")
	cards := newDeck()
	cards.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck to have len 16, but had %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}