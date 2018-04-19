package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ListTasks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ListTasks\n")
}

func CreateTask(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "CreateTasks\n")
}

func ReadTask(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "CreateTasks\n")
}

func UpdateTask(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "CreateTasks\n")
}

func DeleteTask(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "UpdateTasks\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", ListTasks)
	router.POST("/", CreateTask)
	router.GET("/:id", ReadTask)
	router.PUT("/:id", UpdateTask)
	router.DELETE("/:id", DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", router))
}
