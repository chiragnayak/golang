package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	exp_len := int(52)

	if len(d) != exp_len {
		t.Errorf("Mismathc in card count generated, ")
	}
}

func TestSaveDeckToFileReadDeckFromFileRemoveFile(t *testing.T) {

	os.Remove("_test_deck")

	test_deck := newDeck()

	test_deck.saveToFile("_test_deck")

	deckFromFile := readFromFile("_test_deck")

	exp_len := int(52)

	if len(deckFromFile) != exp_len {
		t.Errorf("Mismatch in card count when read from file")
	}

}
