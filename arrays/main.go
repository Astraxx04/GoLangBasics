package main

import "fmt"

// func main() {
// 	var productNames = [4]string{"hi"}
// 	productNames[2] = "hello"
// 	fmt.Println(productNames)
// 	prices := [5]float64{1, 2, 3, 4, 5}
// 	highPrice := prices[1:]
// 	featPrice := highPrice[2:]
// 	fmt.Println(len(featPrice), cap(featPrice))
// }

func main() {
	prices := []float64{1, 2, 3, 4, 5} // slices
	fmt.Println(prices)

	updPrices := append(prices, 6)
	fmt.Println(updPrices)

	newPrices := []float64{6, 7, 8}
	prices = append(prices, newPrices...) // pulls out all elements of the list and provide it as separate elements
	fmt.Println(prices)

	userNames := make([]string, 2, 5) // 2 slots and 5 is the capacity
	userNames[0] = "Red"
	userNames[1] = "Bun"
	userNames = append(userNames, "Julie")
	userNames = append(userNames, "Jobby")
	fmt.Println(userNames)		

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	exercise()
}
