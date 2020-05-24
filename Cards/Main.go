package main

import "fmt"

func main() {
	//var Cards String = "Ace of Spades"
	card := "Ace of Spades"
	card = "Ace of Clubs"  // colon not necessary

	card = newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
