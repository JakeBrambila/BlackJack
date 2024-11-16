package main

import "math/rand"

// Deck represents a standard deck of 52 playing cards.
type Deck struct {
	deck []Card
}

// createDeck initializes the deck with standard playing cards.
// It sets each card with its appropriate suit, number, and royalty rank.
// The deck is ordered in a way where the first 13 cards are Spades,
// the next 13 are Hearts, followed by Diamonds and Clubs.
func (d *Deck) createDeck() {
	d.deck = make([]Card, 52)
	var i, j, k int
	k = 0 // Index for placing cards in the deck

	// Loop over the suits
	for i = 0; i < 4; i++ {
		// Loop over the card values
		for j = 1; j < 14; j++ {
			// Assign specific values based on card number
			if j == 1 {
				d.deck[k].num = Ace
				d.deck[k].royalty = NoRoyalty
			} else if j == 11 {
				d.deck[k].royalty = Jack
				d.deck[k].num = NoCardNum
			} else if j == 12 {
				d.deck[k].royalty = Queen
				d.deck[k].num = NoCardNum
			} else if j == 13 {
				d.deck[k].royalty = King
				d.deck[k].num = NoCardNum
			} else {
				d.deck[k].num = CardNum(j)
				d.deck[k].royalty = NoRoyalty
			}

			// Assign the suit based on the loop index (i)
			if i == 0 {
				d.deck[k].suit = Spade
			} else if i == 1 {
				d.deck[k].suit = Heart
			} else if i == 2 {
				d.deck[k].suit = Diamond
			} else if i == 3 {
				d.deck[k].suit = Club
			}
			k++
		}
	}
}

// shuffleDeck shuffles the cards in the deck randomly.
// It swaps each card with a randomly selected card to ensure the deck is shuffled.
func (d *Deck) shuffleDeck() {
	var i, j int
	var temp Card

	for i = 0; i < 52; i++ {
		j = rand.Intn(52)

		// Swap cards at positions i and j
		temp = d.deck[i]
		d.deck[i] = d.deck[j]
		d.deck[j] = temp
	}
}

// dealCard returns the top card from the deck and removes it from the deck.
func (d *Deck) dealCard() Card {
	var card = d.deck[0]
	return card
}

// removeTopCard removes the top card from the deck.
// is a helper function for deal card
func (d *Deck) removeTopCard() {
	d.deck = d.deck[1:]
}

func (d *Deck) printDeck() {
	for i := 0; i < len(d.deck); i++ {
		d.deck[i].toString()
	}
}
