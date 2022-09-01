package main

import (
	"fmt"
)

func main() {
	fmt.Println("Here we are playing with MAP")
	colors := map[string]string{
		"red":   "#F0003",
		"brown": "#B0040",
	}

	fmt.Println(colors)
	fmt.Println(colors["red"])
	fmt.Println(colors["brown"])
	fmt.Println(colors["missing"]) // prints blank

	var color_var map[string]string // this is just a variable creation.
	// color_var["10"] = "ten" // this will create error
	color_var = make(map[string]string) //assign an object
	fmt.Println(color_var)
	color_var["10"] = "ten"
	fmt.Println(color_var)
	delete(color_var, "10")
	fmt.Println(color_var)
	color_var["1"] = "one"
	color_var["2"] = "two"
	color_var["3"] = "three"
	color_var["4"] = "four"

	printMap(color_var)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
