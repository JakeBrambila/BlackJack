package main

//import "fmt"

func main() {

	card1 := Card{NoSuit, NoCardNum, King}
	card2 := Card{NoSuit, NoCardNum, Jack}
	card3 := Card{NoSuit, NoCardNum, Queen}
	card4 := Card{Spade, Three, NoRoyalty}
	card5 := Card{Heart, Four, NoRoyalty}
	card6 := Card{Diamond, Ten, NoRoyalty}
	cards := []Card{card1, card2, card3, card4, card5, card6}

	PrintCardsHorizontally(cards)
}
