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
	"github.com/joho/godotenv"
)

var db *sql.DB

// must be uppercase for template and pointer for nullable field
type Wine struct {
	Id          int
	Title       string
	Grape       *string
	Origin      *string
	Producer    *string
	Vintage     *int
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

type ContactDetails struct { // issue with types, num saved as string and insecure
	Title       string
	Grape       string
	Origin      string
	Producer    string
	Vintage     int
	Taste       string
	Color       string
	Aroma       string
	Acidity     float64
	Sweetness   float64
	Price       float64
	ISWineScale float64
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
			&wi.ISWineScale) // check err
		Wines = append(Wines, wi)
	}
	if err != nil {
		log.Fatal(err)
	}

	roottmpl := template.Must(template.ParseFiles("root.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := rootPage{
			PageTitle: "Wines",
			Wines:     Wines,
		}
		roottmpl.Execute(w, data)
	})

	formtmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			formtmpl.Execute(w, nil)
			return
		}

		// sanetize/fix input
		// strconv.Atoi can't take empty strings

		func emptyString(string name)  {
			if (r.FormValue(name) = "") {
				r.FormValue(name) = NULL
			}
			return r.FormValue(name)
		}

		Vintage, err := strconv.Atoi()
		if err != nil {
			log.Fatal(err)
		}
		Acidity, err := strconv.Atoi(r.FormValue("acidity"))
		if err != nil {
			log.Fatal(err)
		}
		Sweetness, err := strconv.Atoi(r.FormValue("sweetness"))
		if err != nil {
			log.Fatal(err)
		}
		Price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			log.Fatal(err)
		}
		ISWineScale, err := strconv.Atoi(r.FormValue("isWineScale"))
		if err != nil {
			log.Fatal(err)
		}

		details := ContactDetails{
			Title:       r.FormValue("title"),
			Grape:       r.FormValue("grape"),
			Origin:      r.FormValue("origin"),
			Producer:    r.FormValue("producer"),
			Vintage:     Vintage,
			Taste:       r.FormValue("taste"),
			Color:       r.FormValue("color"),
			Aroma:       r.FormValue("aroma"),
			Acidity:     float64(Acidity),
			Sweetness:   float64(Sweetness),
			Price:       float64(Price),
			ISWineScale: float64(ISWineScale),
		}

		// do something with details
		// Inserts our data into the users table and returns with the result and a possible error.
		// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.

		// result, err := db.Exec(`
		// INSERT INTO users (
		// username,
		// password,
		// created_at
		// ) VALUES (?, ?, ?)`,
		// username,
		// password,
		// createdAt)

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

		id, err := result.LastInsertId()
		fmt.Println(id)

		formtmpl.Execute(w, struct{ Success bool }{true})
	})

	port := "8080"

	fmt.Printf("Listing on port " + port)
	http.ListenAndServe(":"+port, nil)
}
