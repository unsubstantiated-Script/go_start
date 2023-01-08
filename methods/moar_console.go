package methods

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
)

func ConsoleReader() {
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

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program is ending...")

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
}
