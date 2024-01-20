package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	var balanceText = string(data)
	balance, err2 := strconv.ParseFloat(balanceText, 64)

	if err2 != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	var balanceText = fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------------------------")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposite money")
		fmt.Println("3. Withdrawel")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("your deposite: ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)
			accountBalance += depositeAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawel amount: ")
			var withDrawlAmount float64
			fmt.Scan(&withDrawlAmount)

			if withDrawlAmount <= 0 {
				fmt.Println("Invalid amount, must be greater than or 0.")
				continue
			}

			if withDrawlAmount > accountBalance {
				fmt.Println("Invalid amount. you can't withdrawl higher than balance")
				continue
			}

			accountBalance += withDrawlAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
		}
	}
}
