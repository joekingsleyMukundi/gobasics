package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German Sheperd",
	}
	// and now  you can pass th dog as a param here as it has implimented the iterface Animal by using its funtions
	// dog here ha ipmlimented the function says and the function numvber of legs thus matching the requirements for the interface now they impliment the interface
	Printinfo(dog)

	gorilla := Gorilla{
		Name:          "King kong",
		Color:         "black",
		NumberOfTeeth: 24,
	}
	log.Println(gorilla.Color)
	// nowe note that we have not implimmented the saysa nd number of legs funtions in gorilla  so it does not follow the interface thus  it is not of type animal and we now cant use it with the function  print info as it expecte a parameter of type animal
	//TODO: uncommentthis line to see error Printinfo(gorilla)
	// see the erro on uncomment
}

// implimenting the interface functions on Dog struct  to make the dog of type  animal so implimenting interface you just impliment the items it has

func (d Dog) Says() string {
	return "whoo"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

// using the animal type interface
// and now  you can pass th dog as a param here as it has implimented the iterface Animal by using its funtions
func Printinfo(a Animal) {
	log.Println("this animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
