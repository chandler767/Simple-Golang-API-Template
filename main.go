// Created by Chandler Mayo (http://ChandlerMayo.com) and last updated on July 14, 2016.

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"fmt"
	"os"
	"encoding/json"
)

// To use this api: go run main.go
// Or run the makefile and run the version for your os.

func main() {
	r := mux.NewRouter()
	r.Handle("/", StatusHandler).Methods("GET") // Check the status of the API.
	r.Handle("/get/", GetHandler).Methods("GET") // Simple GET request.
	r.Handle("/post/", PostHandler).Methods("POST") // Simple POST request -- accepts form data 'name'.
	r.Handle("/json/{var1}/{var2}/", JSONVarHandler).Methods("GET") // GET request with vars in the url and JSON return.
	http.ListenAndServe(":8091", handlers.LoggingHandler(os.Stdout, r))
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ // http://localhost:8091/  - Check the status of the API.
	w.Write([]byte("This is your API. It is online.\n"))
})

var GetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ // http://localhost:8091/get/  - Simple GET request.
	w.Write([]byte("This is GET request.\n"))
})

// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST -d 'name=Chandler' http://localhost:8091/post/
var PostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ // http://localhost:8091/post/ - Simple POST request -- accepts form data 'name'.
	r.ParseForm()
	name := r.FormValue("name")
	if (len(name) > 0) {
		fmt.Println("The person with name: " + name + " has connected.\n")
		w.Write([]byte("Your name is: " + name + ". This is a post request.\n"))
	} else {
		w.Write([]byte("Missing name.\n"))
	}
})

var JSONVarHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ // http://localhost:8091/json/{var1}/{var2}/ - GET request with vars in the url and JSON return.
	vars := mux.Vars(r) 
	var1 := vars["var1"]
	var2 := vars["var2"]
	jsonreply := map[string]string{"VAR1": var1, "VAR2": var2}
	json.NewEncoder(w).Encode(jsonreply)  
})
