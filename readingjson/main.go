package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"firstname"` //this is to match the data from json to this struct
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func main() {
	myJson := `
	[
		{
			"firstname":"joe",
			"lastname":"khali",
			"phone":"09876754678",
			"email":"joe@gmail.com"
		},
		{
			"firstname":"joe",
			"lastname":"khali",
			"phone":"09876754678",
			"email":"joe@gmail.com"
		}
	]`

	var unmarshaled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshaled)
	if err != nil {
		log.Println("error ", err)
	}

	log.Printf("unmarshaled: %v", unmarshaled)

	// write json from struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "joeking"
	m1.LastName = "mukunsh"
	m1.Email = "jo@gmail.com"
	m1.Phone = "09097656"

	mySlice = append(mySlice, m1)
	var m2 Person
	m2.FirstName = "joeking2"
	m2.LastName = "mukunsh2"
	m2.Email = "jo@gmail.com2"
	m2.Phone = "090976562"

	mySlice = append(mySlice, m2)

	// use json.Marshal in production but use MarshalIndent for dev fo data to be readable
	newJson, err := json.MarshalIndent(mySlice, "", "   ") // the second part is adding a prefix (no need)and the third part is how much spacing do you want so that jsdon spaces enoung to be readabler

	if err != nil {
		log.Println("error: ", err)
	}

	// log.Println(newJson)  this will print bytes
	fmt.Println(string(newJson))
}
