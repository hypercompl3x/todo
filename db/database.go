package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func InitDb() *sql.DB {
	envErr := godotenv.Load()

	if envErr != nil {
		panic(envErr)
	}

	dbUrl := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	if dbUrl == "" || authToken == "" {
		panic("TURSO_DATABASE_URL and TURSO_AUTH_TOKEN must be set")
	}

	fullUrl := dbUrl + "?authToken=" + authToken

	db, err := sql.Open("libsql", fullUrl)

	if err != nil {
		panic(err)
	}

	return db
}