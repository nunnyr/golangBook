package main

import "fmt"

//Creating a new variable is common
// x := "Hello, World"

// func cel(degree int) {
// 	var celsius int
// 	celsius = ((degree - 32) * 5 / 9)
// }

// func main() {
// 	fmt.Print("Enter a number: ")
// 	var input float64
// 	//%f decimal point but no exponent
// 	//.Scanf scans the input texts which is given in the standard input,
// 	//reads it and stores it
// 	fmt.Scanf("%f", &input)
// 	// := is the short assignment statement that can be used in place of a var declaration
// 	// with implicit type.

// 	output := (input - 32) * 5 / 9
// 	fmt.Println(output)
// }
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input * 0.3048)
	fmt.Println(output)
}
