package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/form/v4"
	"github.com/gorilla/mux"
)

// only http errors here

var Decoder = form.NewDecoder()

func FormToWine(r *http.Request) (err error, wine Wine) {
	err = r.ParseForm()
	if err != nil {
		return err, wine
	}

	err = Decoder.Decode(&wine, r.PostForm)
	if err != nil {
		return err, wine
	}

	err = nil
	return err, wine
}

func Roothandler(w http.ResponseWriter, r *http.Request) {
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
		log.Fatal(err)
	}
}

func Addhandler(w http.ResponseWriter, r *http.Request) {
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
}

// no edit = exit:1 nil
func Edithandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Templates/edit.html"))
	Id, err := strconv.ParseInt(mux.Vars(r)["Id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

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

	err, wine := FormToWine(r)
	if err != nil {
		log.Fatal(err)
	}

	err, success := Update(wine, Id)
	if err != nil || !success {
		log.Fatal(err)
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}

func Deletehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Templates/delete.html"))
	Id, err := strconv.ParseInt(mux.Vars(r)["Id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	err, success := Delete(Id)
	if err != nil || !success {
		log.Fatal(err)
	}

	data := struct {
		PageTitle string
		Success   bool
	}{
		PageTitle: "Remove from the winyard",
		Success:   success,
	}

	tmpl.Execute(w, data)
}
