package main

type Hand struct {
	cards []Card
	value int
}

func (h *Hand) addCard(c Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) getHandValue() {
	var total int
	for _, c := range h.cards {
		total += c.getValue()
	}
	h.value = total
}
