package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int {1,2,3,4,5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled) 
	fmt.Println(tripled)

	getTransformerFn1 := getTransformerFn(&numbers)
	moreTransformerFn1 := transformNumbers(&numbers, getTransformerFn1)
	fmt.Println(moreTransformerFn1)

	anonymousMain()
}

func transformNumbers(numbers *[]int, transform transformFn) ([]int) {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFn(nums *[]int) transformFn {
	if(*nums)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(num int) (int) {
	return num*2
}

func triple(num int) (int) {
	return num*3
}

// Can pass functions as paramaters