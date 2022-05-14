package main

import (
	"fmt"
	"gloself/pkg/handlers"
	"net/http"
)

const portNumber = ":8000"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Staring application on port %s", portNumber)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Printf("Error during listening %s", err)
	}
}
