package main

import (
	"log"
	"practice-channel/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)
	intChan <- randomNumber // pass a value into the channel
}

func main() {
	intChan := make(chan int)

	// After "defer", excute as soon as the current function is done.
	// When opening a file or a connection to the database, we don't leave them open forever.
	defer close(intChan)

	// go-routine
	go CalculateValue(intChan)

	// listen for the response to that channel
	// get a value from the channel
	num := <-intChan
	log.Println(num)
}
