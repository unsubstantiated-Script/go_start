package packageone

import "fmt"

//lower case for private capital for public
var privateVar = "I am a private var"
var PublicVar = "I am a public var"

func notExported() {
	fmt.Println("Yeet the non-export")
}

func Exported() {
	fmt.Println("Yeet the export")
}
