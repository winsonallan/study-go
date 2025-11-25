package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, errRevenue := getUserInput("Revenue: ")
	expenses, errExpenses :=  getUserInput("Expenses: ")
	taxRate, errTaxRate := getUserInput("Tax Rate: ")

	if errRevenue != nil || errExpenses != nil || errTaxRate != nil {
		errText := fmt.Sprintf("Error Revenue: %v\nError Expenses: %v\nError Tax Rate: %v", errRevenue, errExpenses, errTaxRate)

		panic(errText)
	}

	ebt, profit, ratio := calcProfits(revenue, expenses, taxRate)
	
	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func storeResults (ebt, profit, ratio float64) {
	ebtText := fmt.Sprintf("EBT: %.1f\n", ebt)
	profitText := fmt.Sprintf("Profit: %.1f\n", profit)
	ratioText := fmt.Sprintf("Ratio: %.3f\n", ratio)
	fileText := fmt.Sprint(ebtText, profitText, ratioText)

	os.WriteFile("profit.txt", []byte(fileText), 0644)
} 

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if (userInput <= 0){
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calcProfits(revenue, expenses, taxRate float64)(float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}