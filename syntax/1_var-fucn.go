package main

import "fmt"

func main() {

	var whatToSay string
	var i int

	whatToSay = "Goobye"
	i = 7

	fmt.Println("Hello, world.")
	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)

	word1, word2 := saySomething2()
	fmt.Println("The function returned", word1, word2)
}

// 1-return
func saySomething() string {
	return "something"
}

// 2-return
func saySomething2() (string, string) {
	return "something", "else"
}
