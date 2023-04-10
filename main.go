package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	Id      int    `json:id`
	Name    string `json:name`
	Content string `json:content`
}

type allTask []task

var tasks = allTask{
	{
		Id:      1,
		Name:    "task one",
		Content: "some content",
	},
}

func homeRoute(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "welcome to my API/home")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeRoute)

	fmt.Println("Initial server on port: :5051")
	log.Fatal(http.ListenAndServe(":5051", router))
}
