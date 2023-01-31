package methods

import "fmt"

//Pointer and interface type
type Animalz interface {
	// Says The functions the interface carries
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Says Methods that go with the above struct
func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

// Says Methods that go with the above struct
func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func TypesOfInterfaces() {

	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	cat := Cat{
		Name:         "cat",
		Sound:        "woem",
		NumberOfLegs: 3,
		HasTail:      true,
	}

	//This currently won't work when no interface is there to define Animal
	Riddle(&cat)
}

//func Riddle(d Dog) {

func Riddle(a Animalz) {
	riddle := fmt.Sprintf(`This animal says "%s" and has "%d" legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
