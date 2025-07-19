package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)

	// error treatment
	if err != nil {
		return 1000, errors.New("failed to find balance.txt file")
	}

	floatValue, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return floatValue, nil // nil = no error
}

func writeBalanceToFile(balance float64) {
	convertedBalace := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(convertedBalace), 0644)
}


func main() {
	var choise int
	var balance, err = readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("------")
		panic(err)
	}

	fmt.Println("Welcome!")


	for {
		fmt.Scan(&choise)

		switch choise {
			case 1:
				fmt.Println("Balance")
				fmt.Println(balance)
			case 2:
				fmt.Print("Deposit: ")
				var deposit float64
				fmt.Scan(&deposit)

				if deposit <= 0 {
					fmt.Println("Invalid operation")
					continue
				}

				balance += deposit
				writeBalanceToFile(balance)
			case 3:
				fmt.Print("Withdraw: ")
				var withdraw float64
				fmt.Scan(&withdraw)

				if withdraw <= 0 || withdraw > balance {
					fmt.Println("Invalid operation")
					continue
				}

				balance -= withdraw
				writeBalanceToFile(balance)
			default:
				fmt.Println("Goodbye")
				return
		}
	}
}