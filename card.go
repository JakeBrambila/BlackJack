package main

import (
	"fmt"
)

type Suit string
type Royalty string
type CardNum int

const (
	Spade   Suit = "♠"
	Heart   Suit = "♥"
	Diamond Suit = "♦"
	Club    Suit = "♣"
	NoSuit  Suit = ""
)

const (
	Jack      Royalty = "J"
	Queen     Royalty = "Q"
	King      Royalty = "K"
	NoRoyalty Royalty = ""
)

const (
	NoCardNum CardNum = 0
	Ace       CardNum = 1
	Two       CardNum = 2
	Three     CardNum = 3
	Four      CardNum = 4
	Five      CardNum = 5
	Six       CardNum = 6
	Seven     CardNum = 7
	Eight     CardNum = 8
	Nine      CardNum = 9
	Ten       CardNum = 10
)

type Card struct {
	suit    Suit
	num     CardNum
	royalty Royalty
}

type QuestionMarkCard struct {
	cardLines []string
}

func (c *QuestionMarkCard) createQuestionMarkCardLines() {
	c.cardLines = make([]string, 5)
	c.cardLines[0] = "+-----+"
	c.cardLines[1] = "|     |"
	c.cardLines[2] = "|  ?  |"
	c.cardLines[3] = "|     |"
	c.cardLines[4] = "+-----+"
}

func (c *Card) getValue() int {
	//royalty
	if c.num == 0 {
		return 10
	}
	return int(c.num)
}

// prints a single card
func (c Card) toString() {
	//no royalty
	if c.royalty == NoRoyalty {
		if c.num < 10 {
			fmt.Printf("+-----+\n")
			fmt.Printf("|%d    |\n", c.num)
			fmt.Printf("|  %s  |\n", c.suit)
			fmt.Printf("|    %d|\n", c.num)
			fmt.Printf("+-----+\n")
			return
		} else if c.num == Ace {
			fmt.Printf("+------+\n")
			fmt.Printf("|%s    |\n", c.suit)
			fmt.Printf("|  A   |\n")
			fmt.Printf("|    %s|\n", c.suit)
			fmt.Printf("+------+\n")
		}
		// 10
		fmt.Printf("+-----+\n")
		fmt.Printf("|%d   |\n", c.num)
		fmt.Printf("|  %s  |\n", c.suit)
		fmt.Printf("|   %d|\n", c.num)
		fmt.Printf("+-----+\n")
	}
	//royalty
	fmt.Printf("+-----+\n")
	fmt.Printf("|%s    |\n", c.suit)
	fmt.Printf("|  %s  |\n", c.royalty)
	fmt.Printf("|    %s|\n", c.suit)
	fmt.Printf("+-----+\n")
}

// returns all the strings of a card in order to use them to be printed horizontally
func (c Card) cardStrings() [5]string {
	var lines [5]string
	// number card
	if c.royalty == NoRoyalty {
		//is Ace
		if c.num == Ace {
			lines[0] = "+-----+"
			lines[1] = fmt.Sprintf("|%s    |", c.suit)
			lines[2] = "|  A  |"
			lines[3] = fmt.Sprintf("|    %s|", c.suit)
			lines[4] = "+-----+"
		} else {
			//not Ace
			if c.num == 10 {
				lines[0] = "+-----+"
				lines[1] = fmt.Sprintf("|%d   |", c.num)
				lines[2] = fmt.Sprintf("|  %s  |", c.suit)
				lines[3] = fmt.Sprintf("|   %d|", c.num)
				lines[4] = "+-----+"
			} else {
				lines[0] = "+-----+"
				lines[1] = fmt.Sprintf("|%d    |", c.num)
				lines[2] = fmt.Sprintf("|  %s  |", c.suit)
				lines[3] = fmt.Sprintf("|    %d|", c.num)
				lines[4] = "+-----+"
			}
		}

	} else {
		// royalty card
		lines[0] = "+-----+"
		lines[1] = fmt.Sprintf("|%s    |", c.suit)
		lines[2] = fmt.Sprintf("|  %s  |", c.royalty)
		lines[3] = fmt.Sprintf("|    %s|", c.suit)
		lines[4] = "+-----+"
	}
	return lines
}

// prints a list of cards in a horizontal format to mimic a cards in a hand
func printCardsHorizontally(cards []Card) {
	//array of string slices
	var linesInCard [5][]string
	fmt.Println("----------------------------------------")

	//adds all the elements of the cards to the linesInCard array
	//e.g. linesInCard[0] = "+-----++-----++-----++-----++-----+"
	//e.g. linesInCard[1] = "|10   ||3    ||4    ||J    ||K    |"
	//e.g. linesInCard[2] = "|  ♠  ||  ♠  ||  ♥  ||  ♦  ||  ♣  |"
	//e.g. linesInCard[3] = "|   10||    3||    4||    J||    K|"
	//e.g. linesInCard[4] = "+-----++-----++-----++-----++-----+"
	for _, card := range cards {
		cardString := card.cardStrings()
		for i := 0; i < 5; i++ {
			linesInCard[i] = append(linesInCard[i], cardString[i])
		}
	}

	for i := 0; i < 5; i++ {
		for _, line := range linesInCard[i] {
			//adds a space to separate the card strings within the same index
			fmt.Print(line + "  ")
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------")
}

func printDealerCards(cards []Card) {
	//array of string slices
	var linesInCard [5][]string
	questionMarkCard := QuestionMarkCard{}
	questionMarkCard.createQuestionMarkCardLines()
	fmt.Println("----------------------------------------")

	//adds all the elements of the cards to the linesInCard array
	//e.g. linesInCard[0] = "+-----++-----+"
	//e.g. linesInCard[1] = "|10   ||3    |"
	//e.g. linesInCard[2] = "|  ♠  ||  ♠  |"
	//e.g. linesInCard[3] = "|   10||    3|"
	//e.g. linesInCard[4] = "+-----++-----+"
	cardString := cards[0].cardStrings()
	for i := 0; i < 5; i++ {
		linesInCard[i] = append(linesInCard[i], cardString[i])
		linesInCard[i] = append(linesInCard[i], questionMarkCard.cardLines[i])
	}

	for i := 0; i < 5; i++ {
		for _, line := range linesInCard[i] {
			//adds a space to separate the card strings within the same index
			fmt.Print(line + "  ")
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------")
}
