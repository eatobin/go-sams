package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}
