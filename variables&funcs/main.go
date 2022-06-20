package main

import (
	"log"
)

func main() {
	var whatToSay string
	var say string
	var secondsay string
	var saySomethingElse string
	// when initializing an iteger with no value  the initial value will be set to 0
	var i int

	whatToSay = saySomething("Hello")

	log.Println(whatToSay)

	// getting the two returned values from the function
	say, secondsay = saySomething1("Goodbye")
	log.Println(say, secondsay)

	// getting only one value from a fuction returning 2 things and ignoring one
	saySomethingElse, _ = saySomething1("Goodbye")
	log.Println(saySomethingElse)

	log.Println(saySomething("finally"))

	i = 7
	i = 8
	log.Println(i)
}

func saySomething(s string) string {
	return s
}

func saySomething1(s string) (string, string) {
	return s, "World"
}
