package methods

import "fmt"

func Expressions() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded)

	//Expression
	ageInTenYears := age + 10

	isATeen := age >= 13

	fmt.Println()
	fmt.Printf("%s is %d years old, but will be %d soon enough. Is it %t that they are a teen?", name, age, ageInTenYears, isATeen)
}
