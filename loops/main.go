package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(1)
	}
	mySlice := []string{"dog", "cat", "gig", "ndovu", "banana", "shoe"}

	// range returns 2  items the index and the item
	for i, x := range mySlice {
		log.Println(i, x)
	}

	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["cat"] = "cat"
	// in a map i willnot bethe index but the key

	for i, x := range myMap {
		log.Println(i, x)
	}
}
