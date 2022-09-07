package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.con",
		"http://golang.org",
		"http://www.amazon.com",
	}

	for _, link := range links {
		if getWebsiteStatus(link){
			fmt.Println(link, " is OK.")
		}else{
			fmt.Println(link, " is NOT RECHEABLE.")
		}
	}
}

func getWebsiteStatus(string link) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	// fmt.Println(resp)
	if resp != nil {
		return true
	}

	retrun false
}
