package my_logger

//A common use case for an infinite loop is to have a channel going to read inputs...
import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func LogerInit() {
	//Read input user 5x and write to log
	reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)

	go ListenForLog(ch)

	fmt.Println("Enter something")

	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Millisecond)
	}
}
