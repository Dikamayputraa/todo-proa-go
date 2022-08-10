package main

import (
	"github.com/Dikamayputraa/todo-proa-go/controllers/taskcontroller"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", taskcontroller.Index)
	mux.HandleFunc("/task/get_form", taskcontroller.GetForm)
	mux.HandleFunc("/task/store", taskcontroller.Store)
	mux.HandleFunc("/task/delete", taskcontroller.Delete)

	log.Println("Starting wen on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
