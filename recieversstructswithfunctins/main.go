package main

import "log"

type myStruct struct {
	FirstName string
}

// function relataed to the strucuct ie its is like a methodso we need to tie the function to mystruct and that is wht we use a pointer in this eg
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct

	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
