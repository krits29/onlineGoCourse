package main

import "testing"

func TestNewDeck(t *testing.T) { //put Test and the name of the function we aer testing
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Unexpected deck length, got %v", len(d))
	}
}
