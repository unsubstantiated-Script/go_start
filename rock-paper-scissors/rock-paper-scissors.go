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

	//Value of computer choice
	computerValue := rand.Intn(2)

	//Reading user input
	reader := bufio.NewReader(os.Stdin)

	//Clearing terminal
	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\r", "", -1)

	if playerChoice == "rock" {
		playerValue = ROCK
	}

	if playerChoice == "paper" {
		playerValue = PAPER
	}

	if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	fmt.Println("Player chose", playerChoice, "and value is", playerValue)
	fmt.Println("Computer value is", computerValue)
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
