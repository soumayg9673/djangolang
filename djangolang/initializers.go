package djangolang

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ReadEnv() {
	log.Print("Reading .env file")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error reading .env file. Error %v", err)
	}
}

func ConnectDb() {
	// read DB_URL connection from .env file
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Error while reading DB_URL")
	}

	log.Print("Trying to connect PostgreSQL DB")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Panicf("Error connecting with PostgeSQL DB. Error %v", err)
	}
	log.Print("Successfully connected to PostgreSQL DB")

	// assign db to Server
	BeServer = ServerConfig{
		DB: db,
	}
}
