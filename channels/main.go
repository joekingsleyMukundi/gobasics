package main

import (
	"log"

	"github.com/joekingsleyMukundi/gobasics/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	// passing a message to the channel

	intChan <- randomNumber
}

func main() {
	// creating a chelnell that can only Holds int
	intChan := make(chan int)

	// close the channel once this program is finished
	defer close(intChan)
	// defer  is used regulory  so this means that whatever comes after the key word defer execute that once the currnt function finishes ie  if connecting to db once connection has been made and everything is okey then closse it  or if writting to a file  closs the connection to the file when don or perfom a certain action when  another is done

	go CalculateValue(intChan)

	// listen to the responce of the channel

	num := <-intChan
	log.Println(num)
}
