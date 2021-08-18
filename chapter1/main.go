//this is known as package declaration and every Go program starts with it.
//packages are Go's way of organizing and reusing code
//Two types of Go programs: executables and libraries

/*
	Executable applications are the kinds of programs that we can run directly from the terminal.
	Libraries are collections of code that we package togetether so that
	we can use them in other programs.
	//edits

*/

/*
	import keyword is how we include code from other packages to use with our program.
	fmt package is short for format. it Implements formatting for input and output.
	Given what we just learned about packages file would contain
	at the top of them.

*/
package main

import "fmt"

//this is a comment
//we assign variables with a var keyword followed with the name
//followed by the type

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

}

// func math() {
// 	fmt.Println(32132 * 42452)
// }
