package main

import "fmt"

/*
asks for: revenue, expenses e tax rate
calculate before tax (EBT), after tax (profit)
calculate ratio (EBT / profit)
*/

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	fmt.Print("EBT: ")
	fmt.Print(ebt)
	fmt.Println("")
	fmt.Print("Profit: ")
	fmt.Print(profit)
	fmt.Println("")
	fmt.Print("Ratio: ")
	fmt.Print(ratio)
}