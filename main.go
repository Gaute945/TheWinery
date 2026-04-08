package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("Templates/root.html"))
		err, wines := ReadAll()
		err, a := ReadOne(1)
		println(*a.Vintage)
		if err != nil {
			log.Fatal(err)
		}

		data := struct {
			PageTitle string
			Wines     []Wine
		}{
			PageTitle: "Wineyard",
			Wines:     wines,
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
	})

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("Templates/add.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		err, wine := FormToWine(r)
		if err != nil {
			log.Fatal(err)
		}

		err, success := Create(wine)
		if err != nil || !success {
			log.Fatal(err)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	// /edit/Id
	r.HandleFunc("/edit/{Id}", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("Templates/edit.html"))
		Id, err := strconv.ParseInt(mux.Vars(r)["Id"], 10, 64)

		if r.Method != http.MethodPost {
			err, wines := ReadOne(Id)
			if err != nil {
				log.Fatal(err)
			}

			data := struct {
				PageTitle string
				Wines     Wine
				Success   bool
			}{
				PageTitle: "Edit the winyard",
				Wines:     wines,
				Success:   false,
			}

			tmpl.Execute(w, data)
			return
		}

		println("post request")

		err, wine := FormToWine(r)
		if err != nil {
			log.Fatal(err)
		}

		success, err := Update(wine)
		if err != nil || !success {
			log.Fatal(err)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	port := "8080"

	fmt.Print("Listing on port " + port)
	http.ListenAndServe(":"+port, r)
}
