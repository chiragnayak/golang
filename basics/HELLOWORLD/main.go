package main

import "fmt"

func main() {
	fmt.Println("Hello World !!")
	cards := []string{"Ace of Jack", "Ten of Spade"}
	fmt.Println("Original List")
	printCards(cards)
	//fmt.Println("append() not reassigned gives below message")
	// append(cards, "Ace of Hearts") // ./main.go:11:2: append(cards, "Ace of Hearts") (value of type []string) is not used
	// printCards(cards)
	fmt.Println("append() and reassigned")
	cards = append(cards, "Ace of Hearts")
	printCards(cards)

	var_declared_not_used()
}

func newCard() string {
	// random cards
	return "Four of Diamonds"
}

func printCards(cards_list []string) {
	for i, card := range cards_list {
		fmt.Println("Card #", i, " > "+card)
	}

}

func var_declared_not_used() {
	for i, num := range []int{1, 2, 3, 4, 5} {
		// fmt.Println(num) // gived compilation error : ./main.go:15:6: i declared but not used
		fmt.Println(i, num)
	}
}
