package methods

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func StringInterpolate() {
	reader = bufio.NewReader(os.Stdin)
	userName := readString("What is your name?")

	age := readInt("How old are you?")
	//fmt.Println("Your name is... "+userName+". You are....", age+"X kakked!")
	fmt.Printf("Your name is %s. You are %dX kakked!\n", userName, age)
}

func promptCursor() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		promptCursor()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Yer kakked! Enter a bleedin name ya savage!")
		} else {
			return userInput
		}
	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		promptCursor()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Yer kakked! Enter a whole number!")
		} else {
			return num
		}
	}
}
