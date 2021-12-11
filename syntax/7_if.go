package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	myNum := 100
	if myNum > 99 && isTrue {
		log.Println("1")
	} else if myNum == 101 || !isTrue {
		log.Println("2")
	} else {
		log.Println("3")
	}
}
