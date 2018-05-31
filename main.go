package main

import "fmt"
import "github.com/eddanger/playing-cards/card"

func main() {
	var suits = []card.Suit{card.SPADE, card.HEART, card.DIAMOND, card.CLUB}
	var ranks = []card.Rank{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}

	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {

			card := card.NewCard(ranks[j], suits[i])
			fmt.Println(card)
		}
	}
}
