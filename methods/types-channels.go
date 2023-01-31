package methods

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

//func TypesOfChannels() {
//	//Firing off a concurrent action
//	go DoSomething("yeet")
//	fmt.Println("This is another thing")
//
//	for {
//		//	Do jack...
//	}
//}

//func DoSomething(s string) {
//	until := 0
//	for {
//		fmt.Println("s is", s)
//		until++
//		if until >= 5 {
//			break
//		}
//	}
//}

//This is a channel and you can only pass one datatype to it
var keyPressChan chan rune

func TypesOfChannels() {
	keyPressChan = make(chan rune)

	//Running this concurrently on its own
	go listenForKeyPress()
	fmt.Println("Press any key, or q to quit")

	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetKey()
		if char == 'q' || char == 'Q' {
			break
		}

		//Send to a channel
		keyPressChan <- char
	}
}

func listenForKeyPress() {
	//This is cheerfully running on its own and listening to receive some sort of information.
	for {
		//Receive a channel
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
