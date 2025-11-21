package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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
	}
	fmt.Println("Connected to db")

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
	fmt.Printf("Listing on port 8080")
	http.ListenAndServe(":8080", nil)

	// r := mux.NewRouter()

	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]

	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// log.Fatal(http.ListenAndServe(":8080", r))
}

// rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", name)
