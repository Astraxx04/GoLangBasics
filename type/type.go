package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}

func type_() {
	var name str = "Max"
	name.log()
}