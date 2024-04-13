package main

import "fmt"

func main() {
	age := 32
	// var agePointer *int
	agePointer := &age
	fmt.Println("Age: ", age)
	// fmt.Println("Adult years: ", getAdultYears(agePointer))
	getAdultYears(agePointer)
	fmt.Println("Age: ", age)
}

func getAdultYears(age *int) {
	*age -= 18
}

// & gives the address
// * gives the value at that address
