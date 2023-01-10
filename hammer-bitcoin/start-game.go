package hammerbitcoin

import "fmt"

func MinerGame() {
	playAgain := true

	for playAgain {
		Play()
		playAgain = GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}
