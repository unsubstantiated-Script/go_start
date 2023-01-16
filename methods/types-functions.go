package methods

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func TypesOfFunctions() {
	z := addTwoNums(2, 4)
	fmt.Println(z)
	fmt.Println(sumMany(4, 5, 6, 7, 8, 9))

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "Woof"
	dog.NumberOfLegs = 44

	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:         "Cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()
}

//Ability to name the return type
func addTwoNums(x, y int) (sum int) {
	//Naked return for very concise functions
	sum = x + y
	return
}

//Ability to take any number of integers

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total
}

//Handling a struct/object as a sort of object method
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}
