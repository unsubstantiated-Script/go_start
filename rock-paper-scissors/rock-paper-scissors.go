package rock_paper_scissors

func RunGame() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	for {
		//Setting round to 1
		game.RoundChan <- 1
		//Wait for a result from this channel
		<-game.RoundChan

		//Stopping the game if gets to round 3
		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			//Need to wait for this before we move on
			<-game.RoundChan
		}
	}

	//fmt.Println("Final score")
	//fmt.Println("==============")
	//fmt.Printf("Player: %d/3, Computer %d/3", playerScore, computerScore)
	//fmt.Println()
	//if playerScore > computerScore {
	//	fmt.Println("Player wins! But really?")
	//} else {
	//	fmt.Println("Computer hosed you!")
	//}
}
