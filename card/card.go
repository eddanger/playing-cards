package card

import "fmt"

type Suit rune

const (
	CLUB    Suit = '♣'
	DIAMOND Suit = '♦'
	HEART   Suit = '♥'
	SPADE   Suit = '♠'
)

type Rank string

const (
	ACE   Rank = "A"
	TWO   Rank = "2"
	THREE Rank = "3"
	FOUR  Rank = "4"
	FIVE  Rank = "5"
	SIX   Rank = "6"
	SEVEN Rank = "7"
	EIGHT Rank = "8"
	NINE  Rank = "9"
	TEN   Rank = "T"
	JACK  Rank = "J"
	QUEEN Rank = "Q"
	KING  Rank = "K"
)

type Card string

func (card Card) String() string {
	return fmt.Sprintf("%s %s of %s", card.unicode(), card.Rank(), card.Suit())
}

func (card Card) unicode() string {
	var suit int32 = 0x0001F0A0
	switch card.Suit() {
	case SPADE:
		suit = 0x0001F0A0
	case HEART:
		suit = 0x0001F0B0
	case DIAMOND:
		suit = 0x0001F0C0
	case CLUB:
		suit = 0x0001F0D0
	}

	return string(suit + int32(card.Rank().Value()))
}

func (card Card) Suit() Suit {
	return []Suit(string(card))[1]
}

func (card Card) Rank() Rank {
	return Rank(card[0])
}

func NewCard(rank Rank, suit Suit) Card {
	return Card(string(rank) + string(suit))
}

func NewCardWithString(card string) Card {
	return Card(card)
}

func (rank Rank) String() string {
	switch rank {
	case ACE:
		return "Ace"
	case TWO:
		return "Two"
	case THREE:
		return "Three"
	case FOUR:
		return "Four"
	case FIVE:
		return "Five"
	case SIX:
		return "Six"
	case SEVEN:
		return "Seven"
	case EIGHT:
		return "Eight"
	case NINE:
		return "Nine"
	case TEN:
		return "Ten"
	case JACK:
		return "Jack"
	case QUEEN:
		return "Queen"
	case KING:
		return "King"
	}
	return ""
}

func (rank Rank) Value() int {
	switch rank {
	case ACE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	case TEN:
		return 10
	case JACK:
		return 11
	case QUEEN:
		return 12
	case KING:
		return 13
	}
	return 0
}

func (suit Suit) String() string {
	switch suit {
	case CLUB:
		return "Clubs"
	case DIAMOND:
		return "Diamonds"
	case HEART:
		return "Hearts"
	case SPADE:
		return "Spades"
	}
	return ""
}
