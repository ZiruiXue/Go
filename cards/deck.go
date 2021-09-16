package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// deck.go - Code that descruobes what a deck is and how it works

// Create a new type of 'deck'
// which is a slice of strings(extends the )
type deck []string

// use d : abbreviation for the type
func (d deck) print() {
	//iterative write cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create and return a list of playing cards. Essentially an array of strings
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Create a 'hand' of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck -> []string -> string -> []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save a list of cards to a file on the local machine
func (d deck) saveToFile(filename string) error {
	// 0666 anyone can read and write this file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Load a list of cards from the local machine
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// error handle
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the roor and entirely quit the program
		fmt.Println("Error:", err)
		// non-zero -> error
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // Ace of Spades, Two of Spades
	return deck(s)
}

// Shuffles all the cards in a deck
// generate a random number
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// problem - every time shuffle in same order
		// why? - because it using the same seeds to generate number
		// how to change -
		// newPosition := rand.Intn(len(d) - 1)

		// use different seed each time
		newPosition := r.Intn(len(d) - 1)
		// easy swap in go
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
