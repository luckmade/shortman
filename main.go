package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	server := http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
	}

	log.Printf("starting server on %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
