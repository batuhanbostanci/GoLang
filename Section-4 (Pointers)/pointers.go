package main

import "fmt"

func main() {

	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Value of age:", *agePointer)

	getAgeToAdultYear(agePointer)
	fmt.Println("Adult years:", age)
}

func getAgeToAdultYear(age *int) {
	*age = *age - 18
}
