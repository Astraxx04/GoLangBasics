package main

import "fmt"

type Product struct {
	name string
	description string
	price float64
}

func exercise() {
	hobbyArr := [3]string{"dance", "sing", "gym"}
	fmt.Println(hobbyArr)
	fmt.Println(hobbyArr[0])
	fmt.Println(hobbyArr[1:])
	mainHobby := hobbyArr[:2]
	fmt.Println(mainHobby)
	fmt.Println(cap(mainHobby))
	mainHobby = mainHobby[1:3] // need to specify the index if u want to extend the selection
	fmt.Println(mainHobby)
	courseGoals := []string{"Learn Go", "Learn basics"}
	fmt.Println(courseGoals)
	courseGoals[1] = "Learn Go in detail"
	courseGoals = append(courseGoals, "Learn basics")
	fmt.Println(courseGoals)
	products := []Product{
		{
			"first-product",
			"First Product",
			12.99,
		},
		{
			"second-product",
			"Second Product",
			125.99,
		},
	}
	fmt.Println(products)
	newProduct := Product{
		"third-product",
		"Third Product",
		199.0,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}