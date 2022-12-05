package main

//Booleans are conditions (if statements of that nature)

import (
	"fmt"
)

//To compare different values we use the (<, >, <=, >=, ==, !=) operators
//The operators check whether the values are true or not therefore giving output to the console "false/true"
//Chained Conditionals
//New Operators (And - &&, Or - ||, Not - !)

func main() {
	x := 8
	y := 4
	val := (true || false) && !!false || y < x
	val2 := (val || !false)
	fmt.Printf("%t", val2 && val)
}
