package main

import "log"

type myStruct struct {
	FirstName string
}

// Receiver can tie myStruct type and below function
// = myStruct type can include function
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	// 1. define myStruct type variable
	var myVar myStruct
	myVar.FirstName = "Taro"

	// 2. define myStruct type variable and set initial value
	myVar2 := myStruct{
		FirstName: "Kenta",
	}

	log.Println(myVar.FirstName)
	log.Println(myVar2.FirstName)

	log.Println(myVar.printFirstName())
	log.Println(myVar2.printFirstName())
}
