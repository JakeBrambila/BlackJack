# Blackjack Game in Go

A simple Blackjack game written in Go, designed to bring the casino experience to your terminal. This project serves as a fun way to practice Go while simulating the classic card game of Blackjack.

## Table of Contents

- [About the Project](#about-the-project)
- [Features](#features)
- [Getting Started](#getting-started)
- [How to Play](#how-to-play)
- [Game Rules](#game-rules)
- [Contributing](#contributing)
- [License](#license)

---

## About the Project

This project is a command-line Blackjack game built in Go. It implements a deck of cards, simulates card dealing, and includes the logic for a basic game of Blackjack. The game displays cards horizontally for easier readability and includes support for standard Blackjack actions like "Hit" and "Stand."

## Features

- **CLI-based Gameplay:** Play Blackjack directly from your terminal.
- **Card Display:** Cards are displayed horizontally for better readability.
- **Basic Blackjack Mechanics:** Supports essential game mechanics like dealing, hitting, and standing.
- **Card Structuring:** Utilizes Go's `struct` and `const` features to represent suits, royalties, and card numbers.
- **Modular Design:** Code is organized to be easy to expand with additional features like betting or multiplayer.

## Getting Started

### Prerequisites

- Go 1.16 or higher is recommended.
- Familiarity with the command line.

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/blackjack-go.git
   cd blackjack-go
   ```
2. **Run the Game:**
   ```bash
   go run main.go
   ```
3. **Build the Executable(optional):**
   ```bash
   go build -o blackjack
    ./blackjack
   ```
## How to Play

Once the game is running, follow the on-screen prompts to play:
	1.	The dealer will deal two cards each to you and themselves.
	2.	You can then choose to:
	•	Hit: Draw an additional card.
	•	Stand: Keep your current hand.
	3.	The goal is to get as close to 21 as possible without going over (busting).
	4.	The dealer will play according to standard Blackjack rules (typically, they must hit until they reach 17 or higher).

## Game Rules

1.	Card Values:
	•	Number cards (2-10) are worth their face value.
	•	Face cards (Jack, Queen, King) are worth 10.
	•	Aces can be worth 1 or 11, depending on which is more favorable to the hand.
2.	Blackjack:
	•	A hand with an Ace and a 10-point card (10, Jack, Queen, or King) is a “Blackjack” and wins the round immediately unless the dealer     also has a Blackjack.
3.	Bust:
	•	If the value of your hand exceeds 21, you “bust” and lose the round.
4.	Dealer Rules:
	•	The dealer must hit until they have at least 17.
5.	Winning Conditions:
	•	You win if your hand’s value is closer to 21 than the dealer’s without exceeding 21.
	•	If both hands have the same value, it’s a “push” (tie).

## Example Output
```
Welcome to Blackjack!
Your hand: +-----+  
           |K    |  
           |  ♠  |
           |    K|    
           +-----+  
Dealer's hand: +-----+  +-----+
               |?    |  |  ♠  |
               |  ♠  |  |    ?|
               +-----+  +-----+

Choose an action:
1. Hit
2. Stand
```

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request with any new features, bug fixes, or improvements.
	1.	Fork the project
	2.	Create your feature branch (git checkout -b feature/new-feature)
	3.	Commit your changes (git commit -m 'Add new feature')
	4.	Push to the branch (git push origin feature/new-feature)
	5.	Open a pull request

Thank you for checking out this repository!
