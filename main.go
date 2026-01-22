package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db *sql.DB

// must be uppercase for template and pointer for nullable field
type Wine struct {
	Id          int64
	Title       string
	Grape       *string
	Origin      *string
	Producer    *string
	Vintage     *int64
	Taste       *string
	Color       *string
	Aroma       *string
	Acidity     *float64
	Sweetness   *float64
	Price       *float64
	ISWineScale float64
}

type rootPage struct {
	PageTitle string
	Wines     []Wine
}

type ContactDetails struct {
	Title       sql.NullString
	Grape       sql.NullString
	Origin      sql.NullString
	Producer    sql.NullString
	Vintage     sql.NullInt64
	Taste       sql.NullString
	Color       sql.NullString
	Aroma       sql.NullString
	Acidity     sql.NullFloat64
	Sweetness   sql.NullFloat64
	Price       sql.NullFloat64
	ISWineScale sql.NullFloat64
}

func isEmptyString(s string) sql.NullString {
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
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBNAME := os.Getenv("DBNAME")

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = DBUSER
	cfg.Passwd = DBPASS
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = DBNAME

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	} else {
		fmt.Println("Connected to db")
	}

	rows, err := db.Query("SELECT Id, Title, Grape, Origin, Producer, Vintage, Taste, Color, Aroma, Acidity, Sweetness, Price, ISwinescale FROM wines")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var Wines []Wine
	for rows.Next() {
		var wi Wine
		rows.Scan(
			&wi.Id,
			&wi.Title,
			&wi.Grape,
			&wi.Origin,
			&wi.Producer,
			&wi.Vintage,
			&wi.Taste,
			&wi.Color,
			&wi.Aroma,
			&wi.Acidity,
			&wi.Sweetness,
			&wi.Price,
			&wi.ISWineScale)
		Wines = append(Wines, wi)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("root.html"))
		data := rootPage{
			PageTitle: "Wines",
			Wines:     Wines,
		}
		tmpl.Execute(w, data)
	})

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("add.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// sanetize/fix input
		details := ContactDetails{}
		details.Title = isEmptyString(r.FormValue("title"))
		details.Grape = isEmptyString(r.FormValue("grape"))
		details.Origin = isEmptyString(r.FormValue("origin"))
		details.Producer = isEmptyString(r.FormValue("producer"))
		details.Vintage = isEmptyInt(r.FormValue("vintage"))
		details.Taste = isEmptyString(r.FormValue("taste"))
		details.Color = isEmptyString(r.FormValue("color"))
		details.Aroma = isEmptyString(r.FormValue("aroma"))
		details.Acidity = isEmptyFloat(r.FormValue("acidity"))
		details.Sweetness = isEmptyFloat(r.FormValue("sweetness"))
		details.Price = isEmptyFloat(r.FormValue("price"))
		details.ISWineScale = isEmptyFloat(r.FormValue("isWineScale"))

		// Inserts our data into the users table and returns with the result and a possible error.
		// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
		result, err := db.Exec(`
		INSERT INTO wines (Title, Grape, Origin, Producer, Vintage, Taste, Color, Aroma, Acidity, Sweetness, Price, ISWineScale
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			details.Title,
			details.Grape,
			details.Origin,
			details.Producer,
			details.Vintage,
			details.Taste,
			details.Color,
			details.Aroma,
			details.Acidity,
			details.Sweetness,
			details.Price,
			details.ISWineScale)
		if err != nil {
			log.Fatal(err)
		}

		Id, err := result.LastInsertId()
		fmt.Println(Id)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	// /edit/Id

	r.HandleFunc("/edit/{Id}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	port := "8080"

	fmt.Print("Listing on port " + port)
	http.ListenAndServe(":"+port, r)
}
