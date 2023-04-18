package main

import (
	"log"
	"net/http"
	

	"github.com/Reticent93/awesomeBanking/internal/repository"
	"github.com/Reticent93/awesomeBanking/internal/repository/dbrepo"
)

//make port a constant
const port = ":4000"

//create a new application struct which holds a couple of application-wide dependencies
type application struct {
	DSN	 string
	DB repository.Repository
	ErrorLog *log.Logger
	InfoLog *log.Logger
}


func main() {

	var app application

	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Banking API"))
	})

	//create a new connection pool
	conn, err := dbrepo.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	
	

	log.Println("Starting the application...")
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

	
}


