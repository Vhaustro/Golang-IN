package main

import (
	"fmt"
)

func main() {

	var number uint64 = 8
	var bl bool
	bl = true
	fmt.Println(bl, number)

	//Print Console & fmt
	//%T prints out the TYPE of the variable e.g (%T, 10) shows 10 is an integer
	//This is used to format strings in the
	//sPrintf allows us to store the output in a variable
	//%v displays the value/statement in default format
	//The percent sign must be in doubles e.g. (%%)
	//Boolea
	//%t will print out true of false value dependin on the value of the variable
	//Numbers can't be formatted as true/false value

	fmt.Printf("Hello %T %v", 10, 50)

	fmt.Println("Welcome to the new server")

	fmt.Printf("Approximate DNS 0880 - localhost")

}

//VARIABLES
//A variable is a method of storing information
//It can be changed/modified but it's just a slot where we place information.
//In a variable the right side is known as the "TYPE"
//Once a variable is declared to have a specific type, it cannot change.
//In a variable there can only be numbers, letters and underscores. It follows the

//Types
//uint8/byte (0=255)(Unsigned Interger) (Positive) (One byte = 8 bits)
//Bit - 0 or 1's to represent information (BASIC)
//Signed integers can be both positive and negative
//Boleans represent true(1bit) or false (0 bit)
//When we declare a variable/anything in golang and it is not used it is consdered a problem/error

//Assignment Expression
//uint8,16,32,64 are explicit because it has been defined and stated what type of variable
//An implicit variable determines or guesses the values within
