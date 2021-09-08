package main

func main() {
	cards := newDeck()
	cards.print()
	// fmt.Println(cards)
}

/*
cards is initialized with invoking the newDeck function




*/

//var card string = "Ace of Spaces"
//^^this is a very explicit way to write this out
// card := "Ace of Spades"
// card = "hearts"
//the below: bc this is a slice of strings, we could just as easily place in a very manually
//created string
//we have declared our new slice of cards

//append does not modify the original array it simply returns a new one
//we assigned the index and the card
