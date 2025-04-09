package database

import (
	"fleet/auth/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	var err error
	DB, err = sqlx.Connect("postgres", config.DB_URL)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connected")
}
