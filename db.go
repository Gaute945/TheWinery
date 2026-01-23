package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func IsEmptyString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func IsEmptyInt(i string) sql.NullInt64 {
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

func IsEmptyFloat(f string) sql.NullFloat64 {
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

func ConnectToDb() (err error) {
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

	return err, wine
}

func ReadAll() (err error, wines []Wine) {
	// db.exec(select * from wines)

	rows, err := db.Query("SELECT Id, Title, Grape, Origin, Producer, Vintage, Taste, Color, Aroma, Acidity, Sweetness, Price, ISwinescale FROM wines")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var w Wine
		rows.Scan(
			&w.Id,
			&w.Title,
			&w.Grape,
			&w.Origin,
			&w.Producer,
			&w.Vintage,
			&w.Taste,
			&w.Color,
			&w.Aroma,
			&w.Acidity,
			&w.Sweetness,
			&w.Price,
			&w.ISWineScale)
		wines = append(wines, w)
	}
	return err, wines
}

func Update(id int64) (err error) {
	// db.exec(insert into wines where id = ?, id)
	// return err

	err = nil
	return err
}

func Delete(id int64) (err error) {
	// db.exec(drop wine where id = ?, id)
	// return err

	return err
}
