package main

import "fmt"

func main() {
	//var Cards String = "Ace of Spades"
	card := "Ace of Spades"
	card = "Ace of Clubs" // colon not necessary

	fmt.Println(card)

	cards := newDeck()

	cards.print()

	fmt.Println(cards)

	hand, remainingCards := deal(cards, 5) //assigning two values

	hand.print()

	remainingCards.print()

}
