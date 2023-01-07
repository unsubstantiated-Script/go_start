package methods

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press Enter when ready."

func MathGame() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	//var answer int

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

}
