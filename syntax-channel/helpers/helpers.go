package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().Unix()) // Giving a different value to the seed, create a random number every execution
	value := rand.Intn(n)
	return value
}
