package main

import (
	"fmt"
	"net/http"

	"github.com/anujyadav001/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
