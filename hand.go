package main

type Hand struct {
	cards []Card
}

func (h *Hand) addCard(c Card) {
	h.cards = append(h.cards, c)
}