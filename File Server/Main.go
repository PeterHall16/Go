package main
import (
  "net/http"
  "strings"
)
func sayHello(w http.ResponseWriter, r *http.Request) { // Create Handler
  message := r.URL.Path // Get the message from the URL
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}
func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}