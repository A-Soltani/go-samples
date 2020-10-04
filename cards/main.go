package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	fiveSpadesCard := getSpadesCard(5)

	fmt.Println(card)
	fmt.Println(fiveSpadesCard)

	setDeck()
}

func getSpadesCard(number int) string {
	return strconv.Itoa(number) + " of Spades"
}

func setDeck() {
	cards := deck{"Ace of Diamonds"} // this is a slice
	cards = append(cards, "6 of Spades")

	cards.print()
}
