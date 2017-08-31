package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Blank Identifier to get webpage
func main() {
	res, _ := http.Get("http://cipherspace.ca")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
