package main

import "fmt"

func anonymousMain() {
	numbers := []int {1,2,3,4,5}

	doubleTrans := createTransformer(2)
	tripleTrans := createTransformer(3)

	transformed := transformNumbersN(&numbers, func(number int)(int){
		return number*2
	}) // anonymous function

	doubledTrans := transformNumbersN(&numbers, doubleTrans)
	tripledTrans := transformNumbersN(&numbers, tripleTrans)

	fmt.Println(transformed)
	fmt.Println(doubledTrans)
	fmt.Println(tripledTrans)
}

func createTransformer(factor int) (func(int)(int)) {
	return func(number int)(int) {
		return number*factor
	}
}

func transformNumbersN(numbers *[]int, transform func(int)(int)) ([]int) {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

// Every anonymous func is a closure