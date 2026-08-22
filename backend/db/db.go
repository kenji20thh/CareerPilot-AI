package db

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")

	}
}
