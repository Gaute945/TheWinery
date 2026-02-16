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

func Create(wine Wine) (err error, success bool) {
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
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if (rowsAffected != 0) && (err == nil) {
		success = true
	} else {
		log.Fatal(err)
	}

	return err, success
}

func ReadOne(id int64) (err error, wine []Wine) {
	row, err := db.Query(`SELECT * FROM wines WHERE Id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var w Wine
		err := row.Scan(
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
		if err != nil {
			log.Fatal(err)
		}

		wine = append(wine, w)
	}

	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}

	return err, wine
}

func ReadAll() (err error, wines []Wine) {
	rows, err := db.Query("SELECT * FROM wines")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var w Wine
		err := rows.Scan(
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
		if err != nil {
			log.Fatal(err)
		}

		wines = append(wines, w)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return err, wines
}

func Update(id int64, changes []string) (err error) {
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
