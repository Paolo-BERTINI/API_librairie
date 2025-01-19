package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Configurations de la database

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:azerty@tcp(localhost:3306)/library"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err, "Erreur de l'ouverture de la database")
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err, "Erreur de ping à la database")
	}
	log.Println("Database connectée avec succès")
}
