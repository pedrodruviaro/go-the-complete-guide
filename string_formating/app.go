package main

import "fmt"

func main() {
	var value float64 = 3.121348172987
	// var value2 float64 = 10.1

	// fmt.Println("The value is:", value)
	// fmt.Printf("The value is: %v", value)

	// fmt.Printf("The values are: %v and %v", value, value2)

	// fmt.Printf("The value is: %.2f", value) // 3.14
	// fmt.Printf("The value is: %.0f", value) // 3

	formattedFv := fmt.Sprintf("The value is: %.2f", value)

	fmt.Print(formattedFv)
}