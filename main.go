package main

import (
	"fmt"
	"html/template"
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

	/*r.HandleFunc("/edit/{Id}", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("Templates/edit.html"))

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
