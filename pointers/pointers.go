package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	fmt.Println("Age address: ", *agePointer)
	fmt.Println("Age: ", age) // 32
	// fmt.Println("Adult years : ", getAdultYears(agePointer))
	editAgeToAdultYears(agePointer)
	fmt.Println("Age: ", age) // 14
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}