package main

import "fmt"

func main() {
	//var card string = "Ace of Spaces"
	//^^this is a very explicit way to write this out
	card := "Ace of Spades"
	card = "hearts"
	fmt.Println(card)

}

/*

the card variable is declared with a type next to it. it means that card
can only ever be that stype.

Go is a statically typed language.

-A dynamically typed language means that we do not care what type
we assign to a variable.
-Go has a few basic types: bool, string, int, float64(a decimal with zeros after)
-with the colon equals, we are relying on the Go compiler to make it equal
to the string
-we only use the := equal when we are assigning a new variable

*/
