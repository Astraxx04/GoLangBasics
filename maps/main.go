package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := []string{"http://google.com", "http://aws.com"}
	websites := map[string]string{
		"google": "http://google.com",
		"amazon": "http://amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	websites["linkedin"] = "http://linkedin.com"
	fmt.Println(websites)
	delete(websites, "google")
	fmt.Println(websites)

	courseRating := make(floatMap, 5)
	courseRating["go"] = 4.9
	courseRating["react"] = 4.8
	courseRating["node"] = 4.6

	courseRating.output()

	for key, value := range courseRating {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
