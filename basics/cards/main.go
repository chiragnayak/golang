package main

import "fmt"

func main() {
	fmt.Println("This is world of Deck of Cards !!")

	fmt.Println("==== FRESH DECK=======")
	new_deck := newDeck()
	new_deck.print()
	fmt.Println("==== GET HAND =======")
	hand, _ := deal(new_deck, 10)
	hand.print()
	fmt.Println("======= HAND AS STRING =====")
	fmt.Println(hand.toString())
	fmt.Println("======= SAVE TO FILE =====")
	hand.saveToFile("card_in_hand")
	fmt.Println("======= READ FROM FILE =====")
	handFromFile := readFromFile("card_in_hand")
	fmt.Println("====  FILE DECK =======")
	handFromFile.print()
	fmt.Println("====  SHFULLE =======")
	shuffledHand := handFromFile.shuffle()
	shuffledHand.print()

	//fmt.Println("==== SHFULLE =======")

}
