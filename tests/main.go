package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(1000.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)

}

func divide(x, y float32) (float32, error) {
	var result float32 // thi will be set initially to 0
	if y == 0 {
		return result, errors.New("connot divide by zero")
	}
	result = x / y
	return result, nil
}
