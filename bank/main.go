package main

import (
	"bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accBalanceFile = "balance.txt"

func main() {
	var accBalance, err = fileops.GetFloatFromFile(accBalanceFile)
	if err != nil {
		fmt.Println("Error: ", err)
		panic("Can't continue")
	}

	fmt.Println("Welcome to Go Bank!!")
	fmt.Println("Reach us at: ", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accBalance)
		case 2:
			fmt.Print("Enter deposit amount: ")
			var depositAmt float64
			fmt.Scan(&depositAmt)
			if depositAmt <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0!")
				continue
			}
			accBalance += depositAmt
			fmt.Println("Balance updated! New Balance is", accBalance)
			fileops.WriteFloatToFile(accBalance, accBalanceFile)
		case 3:
			fmt.Print("Enter withdraw amount: ")
			var withdrawAmt float64
			fmt.Scan(&withdrawAmt)
			if withdrawAmt > accBalance {
				fmt.Println("Invalid amount. You can't withdraw more than what you have!")
				continue
			}
			accBalance -= withdrawAmt
			fmt.Println("Balance updated! New Balance is", accBalance)
			fileops.WriteFloatToFile(accBalance, accBalanceFile)
		default:
			fmt.Println("Thankyou!!")
			return
		}
		fmt.Println("Your choice: ", choice)
	}
}

// Can't use () with for loop
// Just for loop runs indefinitely
// There is no while loop in go everything is done using for loop
// Break is not needed in go lang
// Have to use return to come out of the loop
// We can use _ as a var name
// panic() is an inbuilt method used to show errors

// if(choice == 1) {
// 	fmt.Println("Your Balance is: ", accBalance)
// } else if (choice == 2) {
// 	fmt.Print("Enter deposit amount: ")
// 	var depositAmt float64
// 	fmt.Scan(&depositAmt)
// 	if(depositAmt <= 0) {
// 		fmt.Println("Invalid amount. Must be greater than 0!")
// 		continue;
// 	}
// 	accBalance += depositAmt
// 	fmt.Println("Balance updated! New Balance is", accBalance)
// } else if (choice == 3) {
// 	fmt.Print("Enter withdraw amount: ")
// 	var withdrawAmt float64
// 	fmt.Scan(&withdrawAmt)
// 	if(withdrawAmt > accBalance) {
// 		fmt.Println("Invalid amount. You can't withdraw more than what you have!")
// 		continue;
// 	}
// 	accBalance -= withdrawAmt
// 	fmt.Println("Balance updated! New Balance is", accBalance)
// } else {
// 	fmt.Println("Thankyou!!")
// }
