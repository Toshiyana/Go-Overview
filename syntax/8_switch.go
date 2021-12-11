package main

func main() {
	myVar := "pig"

	switch myVar {
	case "cat":
		print("cat")
	case "dog":
		print("dog")
	case "fish":
		print("fish")
	default:
		print("else")
	}
}
