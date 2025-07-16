package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to GoBank")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check Balance")
	fmt.Println("1. Deposit Money")
	fmt.Println("1. Withdraw Money")
	fmt.Println("4. Exit")

	var choise int
	fmt.Print("Your choice: ")
	fmt.Scan(&choise)

	// wantsCheckBalance := choise == 1

	if choise == 1 {
		fmt.Println("Your balance is:", accountBalance)
	}
}