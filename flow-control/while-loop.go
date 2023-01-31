package flow_control

import (
	"fmt"
	"math/rand"
	"time"
)

func WhileLoop() {
	//Get a random number
	rand.Seed(time.Now().UnixNano())
	i := 1000

	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println("i is ", i)
		if i > 100 {
			fmt.Println("i is ", i, "so loop keep going...")
		}
	}

	fmt.Println("Got", i, "and broke out of loop")
}
