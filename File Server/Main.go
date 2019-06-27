package main

import (
	"fmt"
	"net/http"
	"strings"
)

const port = ":3000"

func sayHello(w http.ResponseWriter, r *http.Request) { // Create Handler
	message := r.URL.Path // Get the message from the URL
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))

}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Println("Running on port " + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

}
