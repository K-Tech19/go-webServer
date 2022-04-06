package main

import (
	"fmt"
	"log"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request){ // "*" start is pointing to the request. THis function usually have two things(parameters) ("w" and "r"). W is your response, while R is our request. 
 
	//checking to see if our request url path is NOT our /hello route then we will shoot an error message 404 status
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound) // shoots error message
		return // returns to running the message
	}
	// Since we don't want to post anthing to our home page we make sure our response method is always "GET"
	if r.Method !="GET" {
		// if its something besides a GET method, an error message is  triggered
		http.Error(w, "method is not supported", http.StatusNotFound)
		return  // returns to running the message
	}


	fmt.Fprintf(w, "hello")
} 


func main(){
	
	// fileServer is a function we get from inside the http package we imported
	fileServer := http.FileServer(http.Dir("./static")) // here we are telling golang to check the static file then it knows to check out the index.html file directly

	// handle function is found inside the http package we imported
	http.Handle("/",fileServer) // starts handling your home route
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)
	
	fmt.Printf("Starting server at port 8000\n") // shows that your server is connected
	if err:= http.ListenAndServe(":8000",nil); err !=nil { // creates the server heart of the software
			log.Fatal(err) // Fatal() func found inside the log import
	}  



}