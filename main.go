package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// fix function to return: "data, err" not "data, err"
	// add http handlers in http.go in form of: r.HandleFunc("/", HomeHandler)

	err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	} else {
		println("Connected to db")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", Roothandler)
	r.HandleFunc("/add", Addhandler).Methods("POST", "GET")
	r.HandleFunc("/edit/{Id}", Edithandler).Methods("POST", "GET")
	r.HandleFunc("/delete/{Id}", Deletehandler).Methods("POST", "GET")

	port := "8080"

	fmt.Print("Listing on port " + port + "\n")
	http.ListenAndServe(":"+port, r)
}
