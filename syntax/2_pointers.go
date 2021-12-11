package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) // &: 値からアドレスを取り出す
	log.Println("After func call, myString is set to", myString)
}

func changeUsingPointer(s *string) {
	// s: pointer
	// *s: value
	newValue := "Red"
	*s = newValue // アドレスsにnewValueの値を代入
}
