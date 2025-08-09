package deck

import (
	"crypto/rand"
	"errors"
	"math/big"
	"github.com/StepPoker/Casino/tree/main/cards/pkg/card"
)

type Deck struct {
	cards []card.Card
}
// NewDeck returns a new ordered 52-card deck.
func NewDeck() *Deck {
	cards := make([]card.Card, 52)
	for i := 0; i < 52; i++ {
		cards[i], _ = card.NewCardFromID(i)
	}
	return &Deck{cards: cards}
}

// Shuffle performs a cryptographically secure shuffle.
func (d *Deck) Shuffle() error {
	n := len(d.cards)
	for i := n - 1; i > 0; i-- {
		j, err := cryptoRandInt(i + 1)
		if err != nil {
			return err
		}
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
	return nil
}

// Draw removes and returns the top card of the deck.
func (d *Deck) Draw() (card.Card, error) {
	if len(d.cards) == 0 {
		return card.Card{}, errors.New("deck is empty")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

// Cards returns a copy of the current cards in the deck.
func (d *Deck) Cards() []card.Card {
	cpy := make([]card.Card, len(d.cards))
	copy(cpy, d.cards)
	return cpy
}

// Remaining returns the number of cards left in the deck.
func (d *Deck) Remaining() int {
	return len(d.cards)
}

// cryptoRandInt returns a secure random int in range [0, max).
func cryptoRandInt(max int) (int, error) {
	bigN := big.NewInt(int64(max))
	n, err := rand.Int(rand.Reader, bigN)
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}
