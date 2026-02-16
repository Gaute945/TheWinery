package main

import (
	"net/http"

	"github.com/go-playground/form/v4"
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
