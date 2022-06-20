package helpers

import (
	"math/rand"
	"time"
)

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {
	// rand.intn takes the length of the random number  ie if you want the random number to be between 0 1nd 10 tou use 10
	// note  when using the same value it result to the same rangom number  ie if you pass 10 always it will give you the same numbr thst is not random so to make it random you do this
	rand.Seed(time.Now().UnixNano())
	// this is called seeding  and it will make this random  numbers
	value := rand.Intn(n)

	return value
}
