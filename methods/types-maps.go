package methods

import "fmt"

//Maps = Key/Value array
//Maps are not sorted. Gotta use the k/v relationships

func TypesOfMaps() {
	intMap := make(map[string]int)

	intMap["three"] = 3
	intMap["two"] = 2
	intMap["one"] = 1

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//remove stuff

	delete(intMap, "one")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//Searching for elements in a map
	el, ok := intMap["four"]

	if ok {
		fmt.Println(el, "is in map true")
	} else {
		fmt.Println("false")
	}

	//Changing values
	intMap["two"] = 4

	for key, value := range intMap {
		fmt.Println(key, value)
	}
}
