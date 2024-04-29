package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/osvauld/demo-create-user/db"
)

var dbPointer *sql.DB

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm alive!")
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUrl := os.Getenv("DB_URL")
	dbPointer = db.CreateDBConnection(dbUrl)

	mux := http.NewServeMux()
	mux.Handle("/", corsMiddleware(http.HandlerFunc(createDemoUserHandler)))

	mux.Handle("/healthCheck", http.HandlerFunc(healthCheckHandler))

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
