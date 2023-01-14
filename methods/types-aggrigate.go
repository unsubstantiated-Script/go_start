package methods

import "log"

//basic types (numbers, strings, bools)
var myInt int //<-probably only use this one
//var myInt16 int16
//var myInt32 int32
//var myInt64 int64
var myUint uint //<-positive values only
var myFloat32 float32
var myFloat64 float64

//For various architecture types

func TypesOfAggrigates() {
	myInt = 10
	myUint = 20
	myFloat32 = 10.1
	myFloat64 = 100.1
	log.Println(myInt, myUint, myFloat32, myFloat64)

	myString := ""
	log.Println(myString)
	myString = "Strings in Go are immutable, you're actually creating a new var here..."
	log.Println(myString)

	//var myBool = true
	////myBool != "Yellow"

}
