package main

import (
	"fmt"
	"net/http"

	"github.com/shah10zeb/go-practice/pkg/handlers"
)

const portNumber = ":8080"



// main application function
func main() {
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/", handlers.Home)
	fmt.Printf("Starting port on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
