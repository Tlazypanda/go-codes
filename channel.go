package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{"http://google.com", "http://facebook.com", "http://youtube.com"}

	//here we are making a channel for communication between main and other go routines
	//since the routines are executing in parallel usually the main routine doesnt wait and exits
	//so to handle this we are adding this block statement through channel so it waits to receive message
	//and hence we are looping over the mesages to be printed
	c := make(chan string)
	for _, website := range websites {
		//go keyword added to make it a go routine to run concurrently hence decreasing wait time
		go checklink(website, c)
	}

	//for pinging sites once
	for i := 0; i < len(websites); i++ {
		//print message from channel
		fmt.Println(<-c)
	}

	//for pinging sites continuously also send  link to c
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checklink(link, c)
		}(l)

	}
}
func checklink(website string, c chan string) {
	res, err := http.Get(website)

	if err != nil {
		fmt.Println("Error", err)
		//message sent to channel
		c <- "down website"
	}

	fmt.Println(website, " is up and running "+res.Status)
	c <- "up and running"
}
