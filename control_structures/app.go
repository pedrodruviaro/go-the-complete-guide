package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to GoBank")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choise int
	fmt.Print("Your choice: ")
	fmt.Scan(&choise)

	// wantsCheckBalance := choise == 1

	if choise == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choise == 2 {
		var moneyToDeposit float64

		fmt.Print("Your deposit: ")
		fmt.Scan(&moneyToDeposit)

		accountBalance += moneyToDeposit

		if moneyToDeposit <= 0 {
			fmt.Println("Invalid operation")
			return
		}

		fmt.Println("Your balance is:", accountBalance)
	} else if choise == 3 {
		var moneyToWithdraw float64

		fmt.Print("Your withdraw: ")
		fmt.Scan(&moneyToWithdraw)

		// validations...
		if moneyToWithdraw <= 0 || moneyToWithdraw > accountBalance {
			fmt.Println("Invalid operation")
			return
		}

		accountBalance -= moneyToWithdraw

		fmt.Println("Your balance is:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}