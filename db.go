package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// must be uppercase for global scope and pointer for nullable field
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

var db *sql.DB

// functions is called by http handler
// http methods is banned

func connectToDb() (err error) {
	err = godotenv.Load()
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

	err = db.Ping()
	return err
}

func Create(wine Wine) (err error) {
	// rows, result := db.exec(insert into wines ? ? ? ? ? ?, Title Grape Origin)
	// for(next(scan(rows)))
	// err = nil
	// return (err)

	result, err := db.Exec(
		`INSERT INTO wines (
		Title, Grape, Origin, Producer, Vintage, 
		Taste, Color, Aroma, Acidity, Sweetness, 
		Price, ISWineScale) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		wine.Title, wine.Grape, wine.Origin,
		wine.Producer, wine.Vintage, wine.Taste,
		wine.Color, wine.Aroma, wine.Acidity,
		wine.Sweetness, wine.Price, wine.ISWineScale,
	)
	print(result.RowsAffected())
	return err
}

func ReadOne(id int64) (err error, wine Wine) {
	// db.exec(select * from wines where id = ?, id)
	// return err, wine
}

func ReadAll() (err error, wines Wine) {
	// db.exec(select * from wines)
	// return err, wines

	rows, err := db.Query("SELECT Id, Title, Grape, Origin, Producer, Vintage, Taste, Color, Aroma, Acidity, Sweetness, Price, ISwinescale FROM wines")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			wines.Id,
			wines.Title,
			&wines.Grape,
			&wines.Origin,
			&wines.Producer,
			&wines.Vintage,
			&wines.Taste,
			&wines.Color,
			&wines.Aroma,
			&wines.Acidity,
			&wines.Sweetness,
			&wines.Price,
			wines.ISWineScale)
		wines = append(wines, wines)
	}
	return err, wines
}

func Update(id int64) (err error) {
	// db.exec(insert into wines where id = ?, id)
	// return err
}

func Delete(id int64) (err error) {
	// db.exec(drop wine where id = ?, id)
	// return err
}
