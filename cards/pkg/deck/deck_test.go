package deck

import (
	"testing"

	"github.com/StepPoker/Casino/cards/pkg/card"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if deck.Remaining() != 52 {
		t.Errorf("expected 52 cards, got %d", deck.Remaining())
	}

	ids := map[int]bool{}
	for _, card := range deck.Cards() {
		id := card.ToID()
		if ids[id] {
			t.Errorf("duplicate card ID: %d", id)
		}
		ids[id] = true
	}
}

func TestShuffle(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	err := deck2.Shuffle()
	if err != nil {
		t.Fatalf("shuffle failed: %v", err)
	}

	// With cryptographic RNG, there is a very low chance this fails accidentally
	sameOrder := true
	for i := 0; i < 52; i++ {
		if deck1.cards[i] != deck2.cards[i] {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Error("deck was not shuffled (still in same order)")
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	initial := deck.Remaining()

	card, err := deck.Draw()
	if err != nil {
		t.Fatalf("unexpected error on draw: %v", err)
	}

	if deck.Remaining() != initial-1 {
		t.Errorf("expected %d cards after draw, got %d", initial-1, deck.Remaining())
	}

	if card.String() == "" {
		t.Error("drawn card is invalid")
	}
}

func TestDrawEmptyDeck(t *testing.T) {
	deck := &Deck{cards: []card.Card{}}
	_, err := deck.Draw()
	if err == nil {
		t.Error("expected error when drawing from empty deck")
	}
}
