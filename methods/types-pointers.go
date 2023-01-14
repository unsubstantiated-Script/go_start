package methods

import "fmt"

//Pointers point to a spot in memory

func TypesOfPointers() {
	//x here holds the value of 10
	x := 10

	//Setting the pointer
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	//Getting and changing the pointer
	*myFirstPointer = 15
	fmt.Println("myFirstPointer is still", myFirstPointer)
	fmt.Println("x is now", x)

	//Need to pass a reference
	changeValueOfPointer(&x)
	fmt.Println("x is finally", x)
}

//Function to change value of pointer
//Argument pointer
func changeValueOfPointer(num *int) {
	//Pointer
	*num = 25
}
