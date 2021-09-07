package main

import "fmt"

type deck []string

//the actual copy of the deck we're working with is
//available in the function as a variable called 'd'

//every variable of type 'deck' can call this function on itself

//below I had put that newDeck returns a string as the type.
//in actuality it returns a new deck, so the type is deck which is declared above
//and it is an array of strings
//we are not adding a receiver to this function

//what we could do with creating this deck is have a slice for the suits
//then create another slice for values

//cards is essentially an object
//cardSuits := []string{"Spades, Hearts, Diamonds"}
//cardValues := []string{"Ace", "Two", "Three"}
//what we can say is
// for each suit in cardSuits
// 	for each value in cardValues
//add a new card of 'value of suit' to the 'cards' deck.


cardSuits := []string{"Spades", "Hearts", "Diamonds"}
cardValyes := []string{"Ace", "Two", "Three"}
func newDeck() deck {
	cards := deck{}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//any variable of type deck now gets access to the "deck" method.
//receivers set up methods on variables that we can create
//the variable d is a reference to the working variable or the instance of the deck
//d is like this or self insisde any other language.
