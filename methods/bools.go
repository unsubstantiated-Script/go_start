package methods

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func Bools() {

	//	apples := 18
	////	oranges := 5
	////
	////	fmt.Println(apples == oranges)
	////	fmt.Println(apples != oranges)
	////	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)

	jack := Employee{
		Name:     "Jacky Jack",
		Age:      27,
		Salary:   4000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      55,
		Salary:   6000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, X := range employees {
		if X.Age > 30 {
			fmt.Println(X.Name, "is 30 or older...")
		} else {
			fmt.Println(X.Name, "is a freaking child...")
		}

		if X.Age > 20 || X.Salary > 10 {
			fmt.Println(X.Name, "makes more than 50000 or is over 10")
		}
	}
}
