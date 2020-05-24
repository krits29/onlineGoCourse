package main

import "fmt"

func main() {
	//var Cards String = "Ace of Spades"
	card := "Ace of Spades"
	card = "Ace of Clubs" // colon not necessary

	card = newCard()

	fmt.Println(card)

	cards := []string{newCard(), newCard()}
	cards = append(cards, "Six of Spades") // add on by recreating the slice

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}
