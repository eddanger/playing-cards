# Playing Cards

Learning Go by creating some playing cards.

## Usage

```go
  card := card.NewCard(card.ACE, card.SPADE)
  fmt.Println(card)

  // 🂡 Ace of Spades
```

```go
  card := card.NewCardWithString("5♦")
  fmt.Println(card)

  // 🃅 Five of Diamonds
```
