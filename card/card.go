package card

import "fmt"

type Card struct {
	rank Rank
	suit Suit
}

func (card Card) String() string {
	return fmt.Sprintf("%s %s of %s", card.unicode(), card.rank, card.suit)
}

func (card Card) unicode() string {
	var suit int32 = 0x0001F0A0
	switch card.suit {
	case SPADE:
		suit = 0x0001F0A0
	case HEART:
		suit = 0x0001F0B0
	case DIAMOND:
		suit = 0x0001F0C0
	case CLUB:
		suit = 0x0001F0D0
	}

	return string(suit + int32(card.rank.Value()))
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank, suit}
}
