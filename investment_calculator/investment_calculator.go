package main

import (
	"fmt"
	"math"
)


func main() {
	const inflationRate float64 = 2.5

	// var investmentAmount, years float64 = 1000, 10
	var investmentAmount float64 // defaults to 0.0
	var years float64 = 10
	var expectedReturnRate float64

	fmt.Print("Investment amount: ")

	// Scan -> read input from console
	fmt.Scan(&investmentAmount) // & -> pointer

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount *  math.Pow((1 + expectedReturnRate / 100), years)
	futureRealValue := futureValue / math.Pow((1 + inflationRate / 100), years)

	fmt.Println("Future Value: $", futureValue)
	fmt.Println("Future Real Value: $", futureRealValue)
}