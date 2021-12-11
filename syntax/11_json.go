package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Taro",
			"last_name": "Sato",
			"age": 20
		},
		{
			"first_name": "Kenta",
			"last_name": "Watabe",
			"age": 15
		}
	]
		`

	// unmarshall : jsonからstringに変換。
	// json.MarshallIndent(): structからjsonに変換。インデントが入るのでMarshall()より見やすい。
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	fmt.Printf("unmershalled: %v\n", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Koji"
	m1.LastName = "Yoshida"
	m1.Age = 18

	var m2 Person
	m2.FirstName = "Risa"
	m2.LastName = "Natsume"
	m2.Age = 25

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("Error marshalling", err)
	}

	fmt.Println(string(newJson))
}
