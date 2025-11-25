package main

import (
	"fmt"

	"example.com/bank-new/fileops"
	"github.com/Pallinder/go-randomdata" // Acquired through using command 'go get github.com/Pallinder/go-randomdata'
)

const accountBalanceFile = "balance.txt"

func main () {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("__________")
		// panic("Can't continue, sorry...")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	
	for {
		presentOptions()
	
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
		// wantsCheckBalance := choice == 1
	
		// switch case version
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
				continue
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for choosing our bank!")
				return
		}
	}
}