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
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644);
}

func main () {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("__________")
		// panic("Can't continue, sorry...")
	}

	fmt.Println("Welcome to Go Bank!")
	
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
		// wantsCheckBalance := choice == 1
	
		// switch case version
		/*
		switch choice {
			case 1:
				fmt.Println("Your balance is", accountBalance)
			case 2:
				var depositAmount float64
				fmt.Print("Your deposit: ")
				fmt.Scan(&depositAmount)
			
				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
			
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
			case 3:
				var withdrawalAmount float64
					
				fmt.Print("Your withdrawal: ")
				fmt.Scan(&withdrawalAmount)
			
				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
			
				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					continue
				}
			
				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount:",accountBalance)
				continue
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for choosing our bank!")
				return
		}
		*/

		// if else version
			if choice == 1 {
				fmt.Println("Your balance is", accountBalance)
			} else if choice == 2 {
				var depositAmount float64
				fmt.Print("Your deposit: ")
				fmt.Scan(&depositAmount)
		
				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
		
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
			} else if choice == 3 {
				var withdrawalAmount float64
				
				fmt.Print("Your withdrawal: ")
				fmt.Scan(&withdrawalAmount)
		
				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
		
				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					continue
				}
		
				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
				continue
			} else {
				fmt.Println("Goodbye!")
				break
			}
	}

	fmt.Println("Thanks for choosing our bank!")
}