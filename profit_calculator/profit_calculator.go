package main

import(
	"fmt"
	"errors"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxRate)

	revenue, err1 := getUserInput("Enter revenue: ")
	if(err1 != nil){
		fmt.Println(err1)
		panic(err1)
	}
	expenses, err2 := getUserInput("Enter expenses: ")
	if(err2 != nil){
		fmt.Println(err2)
		panic(err2)
	}
	taxRate, err3 := getUserInput("Enter tax rate: ")
	if(err3 != nil){
		fmt.Println(err3)
		panic(err3)
	}

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Printf("%.2f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if(userInput <= 0) {
		return 0, errors.New("value must be positive")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}