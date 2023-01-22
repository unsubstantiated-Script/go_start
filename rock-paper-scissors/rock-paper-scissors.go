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

func RunGame() {
	//Setting up game variables
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	playerScore := 0
	computerScore := 0

	//Reading user input
	reader := bufio.NewReader(os.Stdin)

	//Clearing terminal
	clearScreen()

	fmt.Println("Rock, Paper, & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game has three rounds, best of three wins.")
	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Println()
		fmt.Println("Round", i)
		fmt.Println("--------")

		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)
		playerChoice = strings.Replace(playerChoice, "\r", "", -1)

		//Value of computer choice
		computerValue := rand.Intn(2)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			//Reset player value to -1 if invalid choice
			playerValue = -1
		}

		fmt.Println()
		fmt.Println("Player chose", strings.ToUpper(playerChoice))

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

		fmt.Println()

		if playerValue == computerValue {
			fmt.Println("It's a draw")
			//decrement the loop by one so the round repeats on draws...
			i--
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					computerScore = computerWins(computerScore)
				} else {
					playerScore = playerWins(playerScore)
				}
				break
			default:
				fmt.Println("Invalid Choice!")
				i--
			}
		}

		fmt.Println("Final score")
		fmt.Println("==============")
		fmt.Printf("Player: %d/3, Computer %d/3", playerScore, computerScore)
		fmt.Println()
		if playerScore > computerScore {
			fmt.Println("Player wins! But really?")
		} else {
			fmt.Println("Computer hosed you!")
		}
	}
}

func computerWins(score int) int {
	fmt.Println("You loser the computer beat you that round. Do better!")
	return score + 1
}

func playerWins(score int) int {
	fmt.Println("You won that one, but what did you really accomplish here? ")
	return score + 1
}

//Clearing screen based on OS
// clearScreen clears the screen
func clearScreen() {
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
