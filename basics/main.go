package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue , fultureRealValue:= calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println("Future value: ", futureValue)
	fmt.Printf("Future value (adjusted for inflation): %.2f\n", fultureRealValue)

	// formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", fultureRealValue)
	// fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64)(float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64)(fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// } 
// Can also be done like this





// When you use := you can't explicitly set a type
// To declare and set a type use can use var keyword
// Also you can put multiple variable of the same type in a single line
// If you omit the type specification you can put multiple types on the same line
// with const we can't ever change the data value
// You can just declare a value and not assign any value
// You can't declare a var with just a name if you're taking input both type and var should be specified
// Backtick is used if you wanna write a string in many lines