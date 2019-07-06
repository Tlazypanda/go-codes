package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{"http://google.com", "http://facebook.com", "http://youtube.com"}
	for _, website := range websites {
		res, err := http.Get(website)

		if err != nil {
			fmt.Println("Error", err)
		}

		fmt.Println(website, " is up and running "+res.Status)

	}
}
