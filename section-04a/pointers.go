package main

import "fmt"

func main () {
	age := 32 // regular variable

	/* Without Pointers */
	/*
	fmt.Println("Age", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
	*/

	/* With Pointers */
	agePointer := &age
	fmt.Println(agePointer) // This prints the memory address of the pointer variable
	fmt.Println(*agePointer) // This prints the value of the pointer variable

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age int) int {
	return age - 18
}

func editAgeToAdultYears(age *int) {
	*age -= 18
}