package main

import (
	"log"
	"sort"
)

func main() {
	// Pattern 1
	var mySlice []string

	mySlice = append(mySlice, "cat")
	mySlice = append(mySlice, "dog")

	log.Println(mySlice)

	// Pattern 2
	var mySlice2 []int

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 3)
	mySlice2 = append(mySlice2, 1)

	sort.Ints(mySlice2) // sort
	log.Println(mySlice2)

	// Pattern 3
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers[0:2])
}
