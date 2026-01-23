package main

import "net/http"

// sql type mess
func FormToWine(r *http.Request) (err error, wine Wine) {
	// sanetize/fix input
	wine = Wine{}
	wine.Title = IsEmptyString(r.FormValue("title"))
	wine.Grape = IsEmptyString(r.FormValue("grape"))
	wine.Origin = IsEmptyString(r.FormValue("origin"))
	wine.Producer = IsEmptyString(r.FormValue("producer"))
	wine.Vintage = IsEmptyInt(r.FormValue("vintage"))
	wine.Taste = IsEmptyString(r.FormValue("taste"))
	wine.Color = IsEmptyString(r.FormValue("color"))
	wine.Aroma = IsEmptyString(r.FormValue("aroma"))
	wine.Acidity = IsEmptyFloat(r.FormValue("acidity"))
	wine.Sweetness = IsEmptyFloat(r.FormValue("sweetness"))
	wine.Price = IsEmptyFloat(r.FormValue("price"))
	wine.ISWineScale = IsEmptyFloat(r.FormValue("isWineScale"))

	err = nil
	return err, wine
}
