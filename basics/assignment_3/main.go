package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args
	fileName := args[1]
	fmt.Println("File Content :", "\n", readFromFile(fileName))
}

func readFromFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
