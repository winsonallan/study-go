package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	// var investmentAmount, years float64= 1000, 10 //Explicit type declaration

	// expectedReturnRate := 5.5 //Type inference WAY 1 (Convention)
	
	// var expectedReturnRate = 5.5 //Type Inference WAY 2 (Not the convention, use WAY 1 when possible)

	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 //Shorter notation for declaring multiple inferred variables in 1 Line. Not recommended tho as it could make programs harder to read
	
	// investmentAmount := 1000.0
	// years := 10.0
	// expectedReturnRate := 5.5

	var investmentAmount, years float64
	expectedReturnRate := 5.5 // Default value
	
	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	
	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf("Future Value: %.2f\nFuture Value (Adjusted for Inflation): %.2f\n", futureValue, futureRealValue)
	
	// Use backticks to format the output visually
	// fmt.Printf(`
	// Future Value: %.2f
	// Future Value (Adjusted for Inflation): %.2f
	// `, futureValue, futureRealValue) 

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {

	fv := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv := fv /  math.Pow(1 + inflationRate/100, years)

	return fv, rfv
}

/*

func calculateFutureValuesAlternative(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {

	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv /  math.Pow(1 + inflationRate/100, years)

	return fv, rfv
}

*/