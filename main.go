package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// must be uppercase for global scope and pointer for nullable field
type Wine struct {
	Id          int64
	Title       string   `form:"title"`
	Grape       *string  `form:"grape"`
	Origin      *string  `form:"origin"`
	Producer    *string  `form:"producer"`
	Vintage     *int64   `form:"vintage"`
	Taste       *string  `form:"taste"`
	Color       *string  `form:"color"`
	Aroma       *string  `form:"aroma"`
	Acidity     *float64 `form:"acidity"`
	Sweetness   *float64 `form:"sweetness"`
	Price       *float64 `form:"price"`
	ISWineScale float64  `form:"isWineScale"`
}

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
