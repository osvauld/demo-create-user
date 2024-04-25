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

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUrl := os.Getenv("DB_URL")
	dbPointer = db.CreateDBConnection(dbUrl)

	http.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm alive!")
	})

	http.HandleFunc("/createDemoUser", createDemoUserHandler)

	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
