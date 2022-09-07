package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://golang.org",
		"http://www.amazon.com",
	}

	channel := make(chan string)

	// single check each link
	// for _, link := range links {
	// 	go checkLink(link, channel)
	// }
	// check once
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }

	// continuous check each link : for loop variation for infinite loop
	// for _, link := range links {
	// 	go checkLinkLooped(link, channel)
	// }
	// for {
	// 	go checkLinkLooped(<-channel, channel)
	// }

	// or

	// // for loop variation for infinite loop : with receive from channel
	// for _, link := range links {
	// 	go checkLinkLooped(link, channel)
	// }
	// for l := range channel {
	// 	go checkLinkLooped(l, channel)
	// }

	//or
	//using anonymous fuction
	for _, link := range links {
		go checkLinkLooped(link, channel)
	}
	for l := range channel {
		go func(link string) { //to pass value to anonyouns function
			time.Sleep(2 * time.Second)
			checkLinkLooped(link, channel)
		}(l) //value being passed to annoymous function
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is DOWN!", "\n", err)
		channel <- "Looks like " + link + " is DOWN !! (to channel)"
		return
	}
	fmt.Println(link, " is UP!")
	channel <- link + " is UP !! (to channel)"

}

func checkLinkLooped(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is DOWN!", "\n", err)
		channel <- link
		return
	}
	fmt.Println(link, " is UP!")
	channel <- link

}
