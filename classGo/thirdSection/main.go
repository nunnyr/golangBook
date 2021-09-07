package main

import "fmt"

func main() {
	//var card string = "Ace of Spaces"
	//^^this is a very explicit way to write this out
	// card := "Ace of Spades"
	// card = "hearts"
	//the below: bc this is a slice of strings, we could just as easily place in a very manually
	//created string
	//we have declared our new slice of cards

	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

//we informed go that anytime the newCard function is executed, it will return a type string

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
-we can initialize a variable outside of a function, we just can't assign a
value to it

-Go has two basic data structures for handling list of things
-array and slice
-array has a fixed length list of things
-a slice is an array that can grow or shrink
-both must be defined with a data type.
-every element in a slice must be of the same type.

*/
