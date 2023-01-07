package doctor

import (
	"bufio"
	"fmt"
	"go_start/methods"
	"os"
	"strings"
)

func TalkDoc() {

	//Reading the input from the CLI
	reader := bufio.NewReader(os.Stdin)
	//Grabbing Eliza response
	whatToSay := Intro()
	//Print to screen
	methods.SayHelloWorld(whatToSay)

	//Reading user input for when the user hits the enter key looping to keep getting inputs
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		//Clearing input for Windows users to stop program
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		//Clearing input for All other systems users to stop program
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" || userInput == "exit" {
			break
		} else {
			//Testing this by printing user input back...
			methods.SayHelloWorld(Response(userInput))
		}
	}
}
