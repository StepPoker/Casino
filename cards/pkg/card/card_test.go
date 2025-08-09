package card

import (
	"testing"
)

func TestNewCardFromID(t *testing.T) {
	tests := []struct {
		id       int
		expected string
	}{
		{0, "2♣"},
		{12, "A♣"},
		{13, "2♦"},
		{25, "A♦"},
		{26, "2♥"},
		{38, "A♥"},
		{39, "2♠"},
		{51, "A♠"},
	}

	for _, tt := range tests {
		card, err := NewCardFromID(tt.id)
		if err != nil {
			t.Errorf("unexpected error for id %d: %v", tt.id, err)
		}
		if card.String() != tt.expected {
			t.Errorf("card %d = %s, want %s", tt.id, card.String(), tt.expected)
		}
	}
}

func TestCard_ToID(t *testing.T) {
	tests := []struct {
		card     Card
		expected int
	}{
		{Card{Suit: Clubs, Rank: Two}, 0},
		{Card{Suit: Clubs, Rank: Ace}, 12},
		{Card{Suit: Spades, Rank: Ace}, 51},
	}

	for _, tt := range tests {
		id := tt.card.ToID()
		if id != tt.expected {
			t.Errorf("%s.ToID() = %d, want %d", tt.card, id, tt.expected)
		}
	}
}

func TestParseCard(t *testing.T) {
	tests := []struct {
		input    string
		expected Card
	}{
		{"2♣", Card{Suit: Clubs, Rank: Two}},
		{"T♦", Card{Suit: Diamonds, Rank: Ten}},
		{"Q♥", Card{Suit: Hearts, Rank: Queen}},
		{"A♠", Card{Suit: Spades, Rank: Ace}},
	}

	for _, tt := range tests {
		card, err := ParseCard(tt.input)
		if err != nil {
			t.Errorf("unexpected error parsing %q: %v", tt.input, err)
		}
		if card != tt.expected {
			t.Errorf("ParseCard(%q) = %v, want %v", tt.input, card, tt.expected)
		}
	}
}

func TestInvalidInputs(t *testing.T) {
	_, err := NewCardFromID(-1)
	if err == nil {
		t.Error("expected error for negative ID")
	}
	_, err = NewCardFromID(52)
	if err == nil {
		t.Error("expected error for ID > 51")
	}
	_, err = ParseCard("X♠")
	if err == nil {
		t.Error("expected error for invalid rank")
	}
	_, err = ParseCard("A?")
	if err == nil {
		t.Error("expected error for invalid suit")
	}
}
