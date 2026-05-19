package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// functions are called by http handler
// http methods are banned

func ConnectToDb() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = DBUSER
	cfg.Passwd = DBPASS
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "TheWinery"

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	return err
}

// sanetize
// idk what securrity to add here
func Fining(wine Wine) (err error, finedWine Wine) {
	v := reflect.ValueOf(wine)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, value)
	}

	return err, finedWine
}

func Create(wine Wine) (err error, success bool) {
	Fining(wine)
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

func ReadOne(Id int64) (err error, wine Wine) {
	row := db.QueryRow("SELECT * FROM wines WHERE Id = ?", Id)
	if err == sql.ErrNoRows {
		println("ErrNoRows")
		log.Fatal(err)
	}

	err = row.Scan(
		&wine.Id,
		&wine.Title,
		&wine.Grape,
		&wine.Origin,
		&wine.Producer,
		&wine.Vintage,
		&wine.Taste,
		&wine.Color,
		&wine.Aroma,
		&wine.Acidity,
		&wine.Sweetness,
		&wine.Price,
		&wine.ISWineScale)
	if err != nil {
		log.Fatal(err)
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

func Update(wine Wine, Id int64) (err error, success bool) {
	Fining(wine)
	_, err = db.Exec(
		`UPDATE wines SET Id = ?, Title = ?, Grape = ?, Origin = ?, Producer = ?, 
		Vintage = ?, Taste = ?, Color = ?, Aroma = ?, Acidity = ?, 
		Sweetness = ?, Price = ?, ISWineScale = ? WHERE Id = ?`,
		Id, wine.Title, wine.Grape, wine.Origin,
		wine.Producer, wine.Vintage, wine.Taste,
		wine.Color, wine.Aroma, wine.Acidity,
		wine.Sweetness, wine.Price, wine.ISWineScale, Id,
	)
	if err != nil {
		log.Fatal(err)
	} else {
		success = true
	}

	return err, success
}

func Delete(Id int64) (err error, success bool) {
	// db.exec(drop wine where id = ?, id)
	// return err

	_, err = db.Exec(
		`DELETE FROM wines WHERE Id = ?`,
		Id,
	)
	if err != nil {
		log.Fatal(err)
		success = false
	} else {
		success = true
	}

	return err, success
}
