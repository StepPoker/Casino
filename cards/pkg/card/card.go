package card

import (
	"errors"
	"fmt"
)

// Suit constants
type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (s Suit) Symbol() string {
	switch s {
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	case Hearts:
		return "♥"
	case Spades:
		return "♠"
	default:
		return "?"
	}
}

// Rank constants
type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (r Rank) Symbol() string {
	switch r {
	case Ten:
		return "T"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return fmt.Sprintf("%d", r)
	}
}

// Card struct
type Card struct {
	Suit Suit
	Rank Rank
}

// NewCardFromID creates a Card from an integer ID (0–51)
func NewCardFromID(id int) (Card, error) {
	if id < 0 || id > 51 {
		return Card{}, errors.New("card ID must be between 0 and 51")
	}
	suit := Suit(id / 13)
	rank := Rank((id % 13) + 2)
	return Card{Suit: suit, Rank: rank}, nil
}

// ToID returns the integer ID for a Card (0–51)
func (c Card) ToID() int {
	return int(c.Suit)*13 + int(c.Rank) - 2
}

// String returns a 2-character string like "A♠"
func (c Card) String() string {
	return c.Rank.Symbol() + c.Suit.Symbol()
}

// ParseCard parses strings like "A♠", "Q♥", "T♦", "2♣"
func ParseCard(input string) (Card, error) {
	if len([]rune(input)) != 2 {
		return Card{}, errors.New("invalid card format")
	}

	r := string(input[0])
	s := string([]rune(input)[1]) // needed for Unicode

	var rank Rank
	switch r {
	case "2":
		rank = Two
	case "3":
		rank = Three
	case "4":
		rank = Four
	case "5":
		rank = Five
	case "6":
		rank = Six
	case "7":
		rank = Seven
	case "8":
		rank = Eight
	case "9":
		rank = Nine
	case "T":
		rank = Ten
	case "J":
		rank = Jack
	case "Q":
		rank = Queen
	case "K":
		rank = King
	case "A":
		rank = Ace
	default:
		return Card{}, errors.New("invalid rank")
	}

	var suit Suit
	switch s {
	case "♣":
		suit = Clubs
	case "♦":
		suit = Diamonds
	case "♥":
		suit = Hearts
	case "♠":
		suit = Spades
	default:
		return Card{}, errors.New("invalid suit")
	}

	return Card{Suit: suit, Rank: rank}, nil
}
