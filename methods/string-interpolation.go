package methods

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

//Defining our own data type
type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func StringInterpolate() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFavNum("What's your favorite number tool?")
	user.OwnsADog = readDogOwner("Do you own a dog? (y/n)")

	//fmt.Println("Your name is... "+userName+". You are....", age+"X kakked!")

	fmt.Printf("Your name is %s. You are %dX kakked! Owns a dog that %.2fX blows: %t  \n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog,
	)
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

func readFavNum(s string) float64 {
	for {
		fmt.Println(s)
		promptCursor()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Yer kakked! Enter a number!")
		} else {
			return num
		}
	}
}

func readDogOwner(s string) bool {

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please input a y/n response?")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}
	}
}
