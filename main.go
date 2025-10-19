package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Albums struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
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
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	root := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from root page!\n")
	}

	p1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from page1!\n")
	}

	p2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from page2!\n")
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/p1", p1)
	http.HandleFunc("/p2", p2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Albums, error) {
	// An albums slice to hold data from returned rows.
	var albums []Albums

	rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Albums
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
