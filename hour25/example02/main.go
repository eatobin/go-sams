package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome stranger!\n")
}

func Hello(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("nameX"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:nameX", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
