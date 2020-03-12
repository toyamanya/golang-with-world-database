package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty" db:"ID"`
	Name        string `json:"name,omitermpty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty" db:"CountryCode"`
	District    string `json:"district,omitempty" db:"District"`
	Population  int32  `json:"population,omitempty" db:"Population"`
}

type Country struct {
	Code       string `json:"id,omitempty" db:"Code"`
	Name       string `json:"name;omitempty" db:"Name"`
	Population int32  `json:"population,omitempty" db:"Population"`
}

func main() {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")

	searchCityName := os.Args[1]

	city := City{}
	country := Country{}

	// 名前のcityの情報を持ってくる
	db.Get(&city, "SELECT * FROM city WHERE Name='"+searchCityName+"'")
	db.Get(&country, "SELECT Code, country.Name, country.Population FROM country INNER JOIN city ON Code = CountryCode WHERE city.Name = '"+searchCityName+"'")

	rate := float64(city.Population) * float64(100) / float64(country.Population)

	fmt.Printf("%sの人口は%sの内%.2f%%です\n", searchCityName, country.Name, rate)
}
