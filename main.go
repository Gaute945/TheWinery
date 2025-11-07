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

type wine struct {
	// fill from database
	name    string
	vintage int
}

type rootPage struct {
	PageTitle string
	Wines     []wine
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

	roottmpl := template.Must(template.ParseFiles("root.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := rootPage{
			PageTitle: "Wines",
			Wines: []wine{
				{name: "Clos de la Coulée de Serrant", vintage: 2019},
				{name: "Barolo Cannubi", vintage: 2016},
				{name: "Riesling Kabinett", vintage: 2021},
			},
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
