package methods

import "fmt"

//For various architecture types

//Structs
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func TypesOfAggregates() {
	//Arrays
	//Note: Slices are often preferred over Arrays due to more functionality built in
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "bird"

	fmt.Println("First element in array is", myStrings[0])

	//instantiating the struct
	//var myCar Car
	//myCar.NumberOfTires = 4
	//myCar.Luxury = false
	//myCar.BucketSeats = true
	//myCar.Make = "Chevy"
	//myCar.Model = "Pickup"
	//myCar.Year = 1996

	//Another way...
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        false,
		BucketSeats:   true,
		Make:          "Chevy",
		Model:         "Pickup",
		Year:          1996,
	}

	fmt.Printf("My car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)

}
