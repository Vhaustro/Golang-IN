package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Kindly enter your date of birth: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 53)
	fmt.Printf("You will be %d years old at the end 2020", 2020-input)

}

//ParseInt with the StrConv module converts the string (scanner) into an integer variable.
//Therefore it only requires numbers to be entered within the input console
