package methods

import (
	"fmt"
	"sort"
)

//Slices are more functional arrays

func TypesOfSlices() {
	var animals []string

	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "bird")

	fmt.Println(animals)

	//Underscore here is the index of the for loop
	for _, x := range animals {
		fmt.Println(x)
	}

	fmt.Println("Element 0 is...", animals[0])
	//[startPosition:lengthOrsteps]
	fmt.Println("First two elements are...", animals[0:2])
	//Length
	fmt.Println("The slice is this long...", len(animals))
	//Sort
	fmt.Println("Sorted?...", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Sorted now?...", animals)

	animals = deleteFromSlice(animals, 1)

	fmt.Println(animals)
}

func deleteFromSlice(a []string, i int) []string {
	//Finding the last item
	a[i] = a[len(a)-1]
	//Setting it to blank
	a[len(a)-1] = ""
	//Removing from the array
	a = a[:len(a)-1]
	return a
}
