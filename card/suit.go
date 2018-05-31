package card

type Suit rune

const (
	CLUB    Suit = '♣'
	DIAMOND Suit = '♦'
	HEART   Suit = '♥'
	SPADE   Suit = '♠'
)

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
