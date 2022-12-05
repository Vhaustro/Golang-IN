package main

import (
	"fmt"
)

//Number types require to be the same to get an intended value
//Floats and int functions can't be used together
//You can use numeric types and bracket the variables to make them equivalent
//% Mod - gives the remainder after division
//Math Package allows you to obtain a specific mathematical function with in the import section.

func main() {
	var num1 int = 8
	var num2 int = 3
	answer := (num2 % num1) + 5
	fmt.Printf("%d", answer)
}
