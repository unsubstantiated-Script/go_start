package operators

import (
	"fmt"
	"math"
)

func Operate() {
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("area is", area)

	half := 1 / 2
	fmt.Println("half with integer wrong", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("half with integer right", halfFloat)

	badThreeSquared := 3 ^ 2
	fmt.Println("not power of 2", badThreeSquared)
	//Properly squaring
	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("power of 2", goodThreeSquared)

	//Good ol modulus
	remainder := 50 % 3
	fmt.Println("remainder", remainder)

	//unary
	x := 3
	x++
	fmt.Println("incrementing with uniary", x)

	x--
	x--
	fmt.Println("decrementing with uniary", x)
}
