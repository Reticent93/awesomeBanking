package main

import (
	
	"log"
	"net/http"
)

//make port a constant
const port = ":4000"

//create a new application struct which holds a couple of application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Banking API"))
	})

	log.Println("Starting the application...")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

	
}


