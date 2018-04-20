package main

import (
	"net/http"
)

func helloWorld3(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(`{"hello": "world"}`))
}

func main() {
	http.HandleFunc("/", helloWorld3)
	http.ListenAndServe(":8000", nil)
}
