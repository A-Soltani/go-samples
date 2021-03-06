package main

import (
	"strconv"
)

func main() {

	cards := newDeck(10)
	cards.print()

	hands, remainingDeck := deal(cards, 5)

	hands.print()
	remainingDeck.print()

	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// fiveSpadesCard := getSpadesCard(5)

	// fmt.Println(card)
	// fmt.Println(fiveSpadesCard)

	// setDeck()
}

func getSpadesCard(number int) string {
	return strconv.Itoa(number) + " of Spades"
}
