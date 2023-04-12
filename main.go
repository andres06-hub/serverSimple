package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	// fmt.Fprintf(res, "welcome to my API/home ...")
	json.NewEncoder(res).Encode("Welcome to my API/home")
}

func getTasks(res http.ResponseWriter, req *http.Request) {
	fmt.Println("get tasks...")
	//create JSON
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(tasks)
}

func createTask(res http.ResponseWriter, req *http.Request) {
	var newTask task
	//get body
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte("Insert a valid task"))
	}
	//
	json.Unmarshal(reqBody, &newTask)
	//generated id
	newTask.Id = len(tasks) + 1
	//Add task created
	tasks = append(tasks, newTask)
	//Response
	// headers and status code
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newTask)
}

func getTask(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Searching task...")
	// Get data, params
	vars := mux.Vars(req) // Extracts the variables from the request
	// Convert data to number
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(res, "Invalid Id")
		return
	}
	// for each task in the of items
	for _, task := range tasks {
		if task.Id == taskId {
			fmt.Println("Exists task...")
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(task)
			return
		}
	}
	json.NewEncoder(res).Encode("Task not found")
	return
}

func deleteTask(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(res, "Invalid Id")
		return
	}

	for i, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			// json.NewEncoder(res).Encode("")
			fmt.Fprintf(res, "The task with ID %v has been deleted", task.Id)
			// json.NewEncoder(res).Encode("")
			return
		}
	}
	json.NewEncoder(res).Encode("Task not found!")
	return
}

func main() {
	//Created router
	router := mux.NewRouter().StrictSlash(true)
	//Created routers
	router.HandleFunc("/", homeRoute)
	router.HandleFunc("/task", getTasks).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	fmt.Println("Initial server on port: :5051")
	log.Fatal(http.ListenAndServe(":5051", router))
}
