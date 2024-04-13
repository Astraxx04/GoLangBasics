package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 2, 3, 4, 5)
	anotherSum := sumup(numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func factorial(number int)(int) {
	if(number == 0) {
		return 1
	}
	return number * factorial(number-1)
}

func sumup(numbers ...int)(int) {
	sum := 0

	for _, val := range numbers {
		sum += val 
	}
	return sum
}// variadic functions