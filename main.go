package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*	byteslice := make([]byte, 9999)
		res.Body.Read(byteslice)
		fmt.Println(string(byteslice))*/

	//this above code snippet does the same as below copy function - creates a byteslice having enough space 9999
	//and then passing it to read function for it to read the resp body to it and then print the string
	// the below copy function too implements reader- res.Body and writer - os.Stdout

	//io.Copy(os.Stdout, res.Body)
	//the below snippet overrides write method and adds the extra line but essentially prints out the same output

	lw := logwriter{}
	io.Copy(lw, res.Body)
}

func (lw logwriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this ", len(bs))
	return len(bs), nil
}
