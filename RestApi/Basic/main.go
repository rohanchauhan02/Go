package main

import (
	"fmt"
	"net/http"
)
func main() {
	fmt.Println("Web Request")
	// res, err := http.GET(url)
	http.ListenAndServe("localhost:8000",nil) // infine for loop
	//handler is nil here
}
