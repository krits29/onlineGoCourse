package main

import "testing"

func TestNewDeck(t *testing.T) { //put Test and the name of the function we aer testing
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Unexpected deck length, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Six of Hearts" {
		t.Errorf("Expected last card pf Six of Hearts, but got %v", d[len(d) - 1])
	}
}
