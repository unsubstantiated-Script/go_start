package methods

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func KoffeeConsole() {
	err := keyboard.Open()

	//Some error handling if keyboard package won't open for whatever reason
	if err != nil {
		log.Fatal(err)
	}

	//Whatever is opened must be closed!
	//Defer is to delay or ignore error till later after the main function called is finished. Also a way to handle an anonymous or self invoking function.
	defer func() {
		//Ignoring any errors returned if keyboard.Close() decides to break
		_ = keyboard.Close()
	}()

	//Making a map...
	coffees := make(map[int]string)

	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Espresso"
	coffees[6] = "Flat White"

	fmt.Println("MENU")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Espresso")
	fmt.Println("6 - Flat White")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		//Converting letter to int
		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
	}

	fmt.Println("Program is ending...")

}

//reader := bufio.NewReader(os.Stdin)
//
//for {
//	fmt.Print("-> ")
//
//	//Sometimes built in methods return more than one thing. The "_" is just a way of saying to ignore that.
//	userInput, _ := reader.ReadString('\n')
//
//	//latching onto "\n" replacing with an empty string, and doing it everywhere in the method the \n is found
//	userInput = strings.Replace(userInput, "\n", "", -1)
//
//	if userInput == "quit" {
//		break
//	} else {
//		fmt.Println(userInput)
//	}
//}
