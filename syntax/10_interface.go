package main

import "fmt"

// Using the interface, tests will be much easier.

// Inteface definition is lists of function
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Color string
}

type Gorilla struct {
	Name          string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Pochi",
		Color: "Brown",
	}

	fmt.Println(dog.Name, dog.Color)
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Monchicchi",
		NumberOfTeeth: 20,
	}

	fmt.Println(gorilla.Name, gorilla.NumberOfTeeth)
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println(a.Says(), a.NumberOfLegs())
}

// In the best practice of the golang, receiver type is pointer.
// Using pointers, the speed will be faster.
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
