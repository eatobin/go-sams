package main

import (
	"net/http"
)

func helloWorld2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World\n"))
}

func main() {
	http.HandleFunc("/", helloWorld2)
	http.ListenAndServe(":8000", nil)
}
