package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to GoBank")

	var accountBalance float64 = getBalanceFromFile()

	for {
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choise int
		fmt.Print("Your choice: ")
		fmt.Scan(&choise)

		// wantsCheckBalance := choise == 1

		switch choise {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var moneyToDeposit float64

			fmt.Print("Your deposit: ")
			fmt.Scan(&moneyToDeposit)

			accountBalance += moneyToDeposit

			if moneyToDeposit <= 0 {
				fmt.Println("Invalid operation")
				continue
			}

			fmt.Println("Your balance is:", accountBalance)
		case 3:
			var moneyToWithdraw float64

			fmt.Print("Your withdraw: ")
			fmt.Scan(&moneyToWithdraw)

			// validations...
			if moneyToWithdraw <= 0 || moneyToWithdraw > accountBalance {
				fmt.Println("Invalid operation")
				continue
			}

			accountBalance -= moneyToWithdraw

			fmt.Println("Your balance is:", accountBalance)
		default:
			writeBalanceToFile(accountBalance)
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}