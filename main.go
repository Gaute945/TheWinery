package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*func isEmptyString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func isEmptyInt(i string) sql.NullInt64 {
	if i == "" {
		return sql.NullInt64{Valid: false}
	}

	a, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		log.Fatal(err)
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: a, Valid: true}
}

func isEmptyFloat(f string) sql.NullFloat64 {
	if f == "" {
		return sql.NullFloat64{Valid: false}
	}

	a, err := strconv.ParseFloat(f, 64)
	if err != nil {
		log.Fatal(err)
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: a, Valid: true}
}*/

func main() {
	connectToDb()
	if err != nil {
		log.Fatal(err)
	} else {
		println("Connected to db")
	}

	ReadAll()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("root.html"))
		ReadAll()
		tmpl.Execute(w, wines)
	})

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("add.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// http handler
		// sanetize/fix input
		wine := Wine{}
		wine.Title = isEmptyString(r.FormValue("title"))
		wine.Grape = isEmptyString(r.FormValue("grape"))
		wine.Origin = isEmptyString(r.FormValue("origin"))
		wine.Producer = isEmptyString(r.FormValue("producer"))
		wine.Vintage = isEmptyInt(r.FormValue("vintage"))
		wine.Taste = isEmptyString(r.FormValue("taste"))
		wine.Color = isEmptyString(r.FormValue("color"))
		wine.Aroma = isEmptyString(r.FormValue("aroma"))
		wine.Acidity = isEmptyFloat(r.FormValue("acidity"))
		wine.Sweetness = isEmptyFloat(r.FormValue("sweetness"))
		wine.Price = isEmptyFloat(r.FormValue("price"))
		wine.ISWineScale = isEmptyFloat(r.FormValue("isWineScale"))

		Create(wine)
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
