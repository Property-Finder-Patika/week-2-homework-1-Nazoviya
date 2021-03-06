package main

import "fmt"

func main() {
	var (
		name   string
		age    int
		famous bool
	)

	// Example #1
	name = "Newton"
	age = 84
	famous = true

	fmt.Println(name, age, famous)

	// Example #2
	name = "Somebody"
	age = 20
	famous = false

	fmt.Println(name, age, famous)

	// Example #3
	// EXERCISE: Why this doesn't work? Think about it.

	// name = 20
	// name = age

	// save the previous value of the variable
	// to a new variable
	var prevName string = name

	// overwrite the value of the original variable
	// by assigning to it
	name = "Einstein"

	fmt.Println("previous name:", prevName)
	fmt.Println("current name :", name)
}
