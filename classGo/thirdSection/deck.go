package main

import "fmt"

type deck []string

//create and return a list of playing cards. essentially an array of strings.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {

			cards = append(cards, value+" of "+suit)

		}
	}

	return cards

}

//logs out the contents of a deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//there are arguments inside of the function call
//we expect to call this deal function with an existing list of cards
//which we refer to "d deck"
//we are calling it d because its convention
//we are returning two values so we are returning type deck twice

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//any variable of type deck now gets access to the "deck" method.
//receivers set up methods on variables that we can create
//the variable d is a reference to the working variable or the instance of the deck
//d is like this or self insisde any other language.

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
