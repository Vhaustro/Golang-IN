package main

import (
	"fmt"
)

//The condition evaluates the expression to see if it is true/false
//Chained Conditionals
//New Operators (And - &&, Or - ||, Not - !)

func main() {

	name := "Lucas"
	fmt.Println("Welcome Dennis,")
	if name == "Dennis" || name == "Lucas" {
		fmt.Println("Glad to have you back Mr. Ng'ethe")
		fmt.Println("It truly has been a while")
	}
	fmt.Println("It feels as if you identity is incomplete, why's that?")
}
