package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	} else {
		println("Connected to db")
	}

	err, wines := ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("root.html"))
		err, wines = ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, wines)
	})

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("add.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		err, wine := FormToWine(r)
		if err != nil {
			log.Fatal(err)
		}

		err = Create(wine)
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	// /edit/Id

	/*r.HandleFunc("/edit/{Id}", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("edit.html"))

		// validate and sanetize "Id"
		vars := mux.Vars(r)
		Idstring := vars["Id"]
		Id, err := strconv.Atoi(Idstring)
		if err != nil {
			http.Error(w, "Invalid Id", http.StatusBadRequest)
			return
		}

		// calling db for Title
		rows, err := db.Query("SELECT Title FROM wines WHERE Id = ?", Id)
		if err != nil {
			log.Fatal(err)
			return
		}

		defer rows.Close()
		for rows.Next() {
			var Title string
			if err := rows.Scan(&Title); err != nil {
				log.Println("scan error:", err)
				return
			}
			fmt.Print(Title)
		}

		tmpl.Execute(w, vars)
	}) db*/

	port := "8080"

	fmt.Print("Listing on port " + port)
	http.ListenAndServe(":"+port, r)
}
