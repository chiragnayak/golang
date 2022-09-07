package main

import (
	"fmt"
)

/*
Interface

	-- basically a cotract - definition - which any type satisfies will be considered a bot --

	--- to whome it may concern ---

	our program has a new type bot

	if you are of "type" in this program with a function called
	"getGreeting" and you return a string, then you are an honarary
	member of type "bot"

	type <name> interface {
		<function_name> (comman_seperated_arguments)(comman_separated_return_types)
	}

	Questions ..
	-> can a implementing type can have additional functions apart from that required by the interface definition
	-> can a type qualify for multiple interfaces (if it satisfies multiple interfaces contract)
*/
type bot interface {
	getGreeting() string
}

/*
structure definitions
*/
type hindiBot struct {
}

type englishBot struct {
}

type spanishBot struct {
}

/*
Interface to individual function connector of sorts
*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
Individual implementation -> same name, different receiver functions
*/

func (hindiBot) getGreeting() string {
	//some custom logic
	return "Namaskaram !!"
}

func (eb englishBot) getGreeting() string { //if you do not want to use the "eb" value, you can remove variable name, just keep the reciever type
	//some custom logic
	return "Hi There !!"
}

func (sb spanishBot) getGreeting() string {
	//some custom logic
	return "Hola !!"
}

func main() {

	hb := hindiBot{}
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(hb)
	printGreeting(eb)
	printGreeting(sb)

}
