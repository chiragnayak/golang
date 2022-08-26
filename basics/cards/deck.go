package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

type deck []string // to be read as slice of string

// receiver
// any varuable of type receiver_type will get access to fucntion func_name below
// receiver variable : actual copy of the receiver type we are working with
// func (receiver_variable recevier_type) func_name() return_type {
//
//}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	numbers := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	symbols := []string{"Heart", "Spade", "Club", "Diamond"}
	newDeck := deck{} // creates a new variable of type deck

	for _, s := range symbols {
		for _, n := range numbers {
			newDeck = append(newDeck, n+" of "+s)
		}
	}

	return newDeck
}

/*
deal()
should be called upon shuffeled cards
to take input on how many cards to deal from shuffled deck
fetch x number of cards from shuffeked deck to new deck and return
from shuffeled cards, cards that are given in hand (deal) must be removed

slicing using [start:end] syntax

slice_name[start_index_including : upto_not_including]

slice_name[1:3]  // from 1 till 2
slice_name[:3] // from start till 2
slice_name[3:] // from 3 till ??
*/
func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingDeck := d[handSize:]

	return hand, remainingDeck
}

// func (d deck) toString() string {
// 	asString := string("")
// 	for i, card := range d {
// 		if i != len(d)-1 {
// 			asString = asString + card + ","
// 		} else {
// 			asString = asString + card
// 		}
// 	}

// 	return asString
// }

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) {
	// convert deck to string
	deckStr := d.toString()

	// convert string to byte array
	deckByte := []byte(deckStr)

	// fmt.Println(deckByte)

	// write to file
	err := ioutil.WriteFile(fileName, deckByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readFromFile(fileName string) (d deck) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	deckStr := string(content)
	deckSlice := strings.Split(deckStr, ",")
	return deckSlice
}

func (d deck) shuffle() deck {
	limit := len(d) - 1

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i, _ := range d {
		currPos := i
		newPos := r.Intn(limit) //this will give exactly same position every time the loop is complete as we are using same seed (default)
		//to fix this you need to use new seed everytime you need to generate a random number
		fmt.Println(currPos, newPos)
		// temp := d[currPos]
		// d[currPos] = d[newPos]
		// d[newPos] = temp
		d[currPos], d[newPos] = d[newPos], d[currPos]

	}

	return d
}
