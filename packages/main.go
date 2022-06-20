package main

import (
	"log"

	"github.com/joekingsleyMukundi/gobasics/helpers"
)

func main() {
	log.Println("hello world")

	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
