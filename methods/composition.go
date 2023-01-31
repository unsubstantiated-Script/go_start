package methods

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car1 struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	//Embedding the vehichle here instead of inheritance
	Vehicle Vehicle
}

//This function is tied to the Vehicle struct
func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers", v.NumberOfPassengers)
	fmt.Println("Number of wheels", v.NumberOfWheels)
}

func (c Car1) showCar1() {
	fmt.Println("Make", c.Make)
	fmt.Println("Model", c.Model)
	fmt.Println("Year", c.Year)
	fmt.Println("Is Electric?", c.IsElectric)
	fmt.Println("Is Hybrid?", c.IsHybrid)

	//	Can also invoke the vehicle function here...
	c.Vehicle.showDetails()
}

func Compositions() {
	pickup := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	chevyPickup := Car1{
		Make:       "Chevy",
		Model:      "Pickup",
		Year:       1996,
		IsElectric: false,
		IsHybrid:   false,
		Vehicle:    pickup,
	}

	chevyPickup.showCar1()

	fmt.Println()

	teslaX := Car1{
		Make:       "Tesla",
		Model:      "X",
		Year:       2020,
		IsElectric: true,
		IsHybrid:   true,
		Vehicle:    pickup,
	}

	teslaX.showCar1()
}
