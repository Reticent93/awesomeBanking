package main

import "net/http"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//write if statement to check if the current request URL path exactly matches "/".
	if r.URL.Path != "/" {
		// If it doesn't, use the w.WriteHeader() method to send a 404 status code and the w.Write() method to write a "Not Found" response body. We then return from the handler so that the subsequent code is not executed.
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return	
	}
	// Use the w.Write() method to write a byte slice containing "Hello from Banking" as the response body.
	w.Write([]byte("Hello from Banking API"))
}