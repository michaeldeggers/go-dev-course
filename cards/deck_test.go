package main

import (
	"os"
	"testing"
)

// t param required
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first item to be 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first item to be 'King of Clubs', but got '%v'", d[51])
	}
}

// ignoring errors
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	_ = os.Remove("_decktesting")

	d := newDeck()
	_ = d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")
	if len(ld) != 52 {
		t.Errorf("Expected deck of length 52, got %v", len(ld))
	}

	_ = os.Remove("_decktesting")
}
