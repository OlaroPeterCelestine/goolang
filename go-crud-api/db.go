package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func initDB() {
	url := os.Getenv("DATABASE_URL") // ← we'll set this in Render
	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("❌ DB Open Error:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("❌ DB Ping Error:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}