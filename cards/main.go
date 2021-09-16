package main

// main.go - Code to create and manipulate a deck
// deck_test.go - Code to autonomatically test the deck

func main() {
	// var card string = "Ace of Spades"

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

}

// first time use :=
// each records have same identity in array,slice
// Array: Fixed length list of things
// Slice: An array that can grow or shrink
// slice[startIndexIncluding:upToNotIncluding]
// Type conversion: Type we want(value we have) eg. []byte("Hi")
