package main

import "log"

func main() {
	// Example 1
	for i := 0; i < 5; i++ {
		log.Println(i)
	}

	// Example 2
	// Slices
	mySlices := []string{"a", "b", "c", "d", "e"}

	for index, value := range mySlices {
		log.Println(index, value)
	}

	for _, value := range mySlices {
		log.Println(value)
	}

	// Example 3
	// Map
	myMap := make(map[string]string)
	myMap["cat"] = "Nyan"
	myMap["dog"] = "Pochi"

	for key, value := range myMap {
		log.Println(key, value)
	}

	// Example 4
	// String
	myString := "My name is Taro"

	for index, byteOfSlices := range myString {
		log.Println(index, ":", byteOfSlices)
	}

	// Example 5
	// Custom type
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Taro", "Sato", "taro@gmail.com", 15})
	users = append(users, User{"Ken", "Watabe", "ken@gmail.com", 20})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}

}
