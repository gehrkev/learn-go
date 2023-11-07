package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[0])
	}
}

func TestCompareDecks(t *testing.T) {
	d1 := newDeck()
	d2 := newDeck()

	if !compareDecks(d1, d2) {
		t.Errorf("compareDecks returned false for equal decks")
	}

	d3 := newDeck()
	d4 := newDeck()
	d4[0] = "Now it's different"

	if compareDecks(d3, d4) {
		t.Errorf("compareDecks returned true for different decks")
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	if !compareDecks(loadedDeck, d) {
		t.Errorf("Loaded deck is different from saved deck")

		for i, card := range loadedDeck {
			if card != d[i] {
				fmt.Printf("Index %d - Loaded: %s, Saved: %s\n", i, loadedDeck[i], d[i])
			}
		}
	}

	os.Remove("_decktesting")
}
