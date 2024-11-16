package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//import "fmt"

func main() {
	fmt.Println("Welcome to Blackjack!")
	fmt.Println()
	for {
		fmt.Println("1. Start game")
		fmt.Println("2. Exit")
		choice, err := ReadIntegerInput("Enter your choice: ")
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}
		if choice == 1 {
			fmt.Println()
			fmt.Println("Starting game...")
			time.Sleep(time.Second)
			fmt.Println()
			playBlackjack()
		} else if choice == 2 {
			fmt.Println()
			fmt.Println("Exiting game...")
			break
		}
	}

}

func playBlackjack() {
	deck := Deck{}
	deck.createDeck()
	deck.shuffleDeck()

	dealerHand := Hand{}
	playerHand := Hand{}

	playerHand.addCard(deck.dealCard())
	deck.removeTopCard()
	playerHand.addCard(deck.dealCard())
	deck.removeTopCard()

	dealerHand.addCard(deck.dealCard())
	deck.removeTopCard()
	dealerHand.addCard(deck.dealCard())
	deck.removeTopCard()

	for {
		clearConsole()
		fmt.Println("Player's hand:")
		playerHand.getHandValue()
		fmt.Printf("Total: %d\n", playerHand.value)
		printCardsHorizontally(playerHand.cards)

		fmt.Println("Dealer's hand:")
		dealerHand.getHandValue()
		printDealerCards(dealerHand.cards)

		fmt.Println()
		fmt.Println("1. Hit")
		fmt.Println("2. Stay")
		choice, err := ReadIntegerInput("Enter your choice: ")
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}

		if choice == 1 {
			playerHand.addCard(deck.dealCard())
			deck.removeTopCard()
			playerHand.getHandValue()
			if playerHand.value == 21 {
				fmt.Println()
				fmt.Println("Blackjack! You win.")
				fmt.Println("Player's hand:")
				printCardsHorizontally(playerHand.cards)
				fmt.Println("Dealer's hand:")
				printCardsHorizontally(dealerHand.cards)
				fmt.Println()
				break
			}

			fmt.Println("Player's hand:")
			fmt.Printf("Total: %d\n", playerHand.value)
			printCardsHorizontally(playerHand.cards)

			fmt.Println("Dealer's hand:")
			printDealerCards(dealerHand.cards)

			if youreBusting(playerHand.value) {
				clearConsole()
				fmt.Println("You busted! Dealer wins.")
				fmt.Println("Player's hand:")
				printCardsHorizontally(playerHand.cards)
				fmt.Println("Dealer's final hand:")
				printCardsHorizontally(dealerHand.cards)
				break
			}
		} else if choice == 2 {
			dealerHand.getHandValue()
			fmt.Println("\nDealer is now drawing a card...")
			time.Sleep(time.Second)
			dealerHand.addCard(deck.dealCard())
			deck.removeTopCard()
			dealerHand.getHandValue()

			//dealer has to keep hitting if his hand is less than 17
			for dealerHand.value < 17 {
				dealerHand.addCard(deck.dealCard())
				deck.removeTopCard()
				fmt.Println("\nDealer is now drawing again...")
				time.Sleep(time.Second)
				dealerHand.getHandValue()
				// fmt.Println("Dealer's hand:")
				// printDealerCards(dealerHand.cards)
				fmt.Println()
			}

			if youreBusting(dealerHand.value) {
				clearConsole()
				fmt.Println("Dealer busted! You win.")
				fmt.Println("Player's hand:")
				printCardsHorizontally(playerHand.cards)
				fmt.Println("Dealer's final hand:")
				printCardsHorizontally(dealerHand.cards)
				fmt.Println()
				break
			}

			if dealerHand.value > playerHand.value {
				clearConsole()
				fmt.Println("Dealer wins.")
				fmt.Println("Player's hand:")
				printCardsHorizontally(playerHand.cards)
				fmt.Println("Dealer's final hand:")
				printCardsHorizontally(dealerHand.cards)
				fmt.Println()
				break
			}

			if dealerHand.value < playerHand.value {
				clearConsole()
				fmt.Println("You win.")
				fmt.Println("Player's hand:")
				printCardsHorizontally(playerHand.cards)
				fmt.Println("Dealer's final hand:")
				printCardsHorizontally(dealerHand.cards)
				fmt.Println()
				break
			}

			if dealerHand.value == playerHand.value {
				clearConsole()
				fmt.Println("It's a tie!")
				fmt.Println("Player's hand:")
				printCardsHorizontally(playerHand.cards)
				fmt.Println("Dealer's final hand:")
				printCardsHorizontally(dealerHand.cards)
				fmt.Println()
				break
			}
		}
	}
}

func youreBusting(handValue int) bool {
	if handValue > 21 {
		return true
	}
	return false
}

func ReadIntegerInput(prompt string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin": // macOS
		cmd := exec.Command("clear") // macOS clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux": // Linux
		fmt.Print("\033[H\033[2J") // ANSI escape sequence for clearing the terminal
	default: // Fallback for other Unix-based systems
		fmt.Print("\033[H\033[2J") // ANSI escape sequence for clearing the terminal
	}
}
