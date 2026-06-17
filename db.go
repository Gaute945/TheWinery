package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// must be uppercase for global scope and pointer for nullable field
type Wine struct {
	Id          int64
	Title       string          `form:"title"`
	Grape       sql.NullString  `form:"grape"`
	Origin      sql.NullString  `form:"origin"`
	Producer    sql.NullString  `form:"producer"`
	Vintage     sql.NullInt64   `form:"vintage"`
	Taste       sql.NullString  `form:"taste"`
	Color       sql.NullString  `form:"color"`
	Aroma       sql.NullString  `form:"aroma"`
	Acidity     sql.NullFloat64 `form:"acidity"`
	Sweetness   sql.NullFloat64 `form:"sweetness"`
	Price       sql.NullFloat64 `form:"price"`
	ISWineScale float64         `form:"isWineScale"`
}

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

func Fining(wine Wine) Wine {
	v := reflect.ValueOf(&wine).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() != reflect.Struct {
			continue
		}

		valid := f.FieldByName("Valid")
		if !valid.IsValid() || !valid.CanSet() {
			continue
		}

		inner := f.Field(0)
		if !inner.IsZero() {
			valid.SetBool(true)
		}
	}

	return wine
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

	wine = Fining(wine)

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
