package rock_paper_scissors

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

//Tied to the Game type
func (g *Game) Rounds() {
	//Using to select and process input for the game
	//Print to screen
	//Keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			//Incrementing the round number
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}
	}
}

//Clearing screen based on OS
// clearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	fmt.Println("Rock, Paper, & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game has three rounds, best of three wins.")
	fmt.Println()
}

func (g *Game) PlayRound() bool {
	//Seeding random number
	rand.Seed(time.Now().UnixNano())

	//Setting player's value
	playerValue := -1

	//Printing prompt for player
	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("______")
	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\r", "", -1)

	//Value of computer choice
	computerValue := rand.Intn(3)

	//Processing input
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose Rock")
		break
	case PAPER:
		fmt.Println("Computer chose Paper")
		break
	case SCISSORS:
		fmt.Println("Computer chose Scissors")
		break
	default:
	}

	//Printing out player choice via a channel we've setup in the other file
	fmt.Println()

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!"
		//decrement the loop by one so the round repeats on draws...ends round
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid Choice!"
			return false
		}
	}

	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "You loser the computer beat you that round. Do better!"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "You won that one, but what did you really accomplish here? "
}
