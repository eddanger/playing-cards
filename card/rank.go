package card

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

func (rank Rank) Index() int {
	ranks := []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}

	for i, _ := range ranks {
		if ranks[i] == rank {
			return i
		}
	}
	return -1
}

func (rank Rank) Value() int {
	return rank.Index() + 1
}

func (rank Rank) String() string {
	strings := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	return strings[rank.Index()]
}
