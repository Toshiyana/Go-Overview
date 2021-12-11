package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Pattern1
	// Declare map
	myMap := make(map[string]string)

	myMap["dog"] = "Taro"
	myMap["cat"] = "Nyan"

	log.Println(myMap["dog"]) // Taro
	log.Println(myMap["cat"]) // Nyan

	myMap["dog"] = "Pochi"    // change value
	log.Println(myMap["dog"]) // "Pochi"

	// Pattern2
	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	log.Println(myMap2["First"])

	// Pattern3
	myMap3 := make(map[string]User)
	me := User{
		FirstName: "Kenta",
		LastName:  "Sato",
	}

	myMap3["me"] = me
	log.Println(myMap3["me"].FirstName)

}
