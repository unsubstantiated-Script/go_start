package methods

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press Enter when ready."

func MathGame() {
	//Seed to avoid getting the same number every time. Based upon finite time changes which always change.
	rand.Seed(time.Now().UnixNano())

	//Adding 2 here keeps us in bounds to avoid bad division operations
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	reader := bufio.NewReader(os.Stdin)
	// Display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1-10 and press ENTER when ready")

	reader.ReadString('\n')

	// Take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// Give them the answer
	answer = firstNumber*secondNumber - subtraction

	fmt.Println("The answer is...", answer, "Buh-Bye!")
}
