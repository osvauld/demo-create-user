package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/osvauld/demo-create-user/db"

	"golang.org/x/crypto/bcrypt"
)

func createRandomUser() (string, string) {
	gofakeit.Seed(0)
	username := gofakeit.Username()
	tempPassword := gofakeit.Password(true, true, true, true, true, 12)

	return username, tempPassword
}

func HashPassword(password string) (string, error) {
	// The second argument is the cost of hashing, which determines how much time is needed to calculate the hash.
	// The higher the cost, the more secure the hash, but the longer it will take to generate.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func createDemoUserHandler(w http.ResponseWriter, r *http.Request) {

	username, tempPassword := createRandomUser()

	hashedPassword, err := HashPassword(tempPassword)
	if err != nil {
		log.Println("Error hashing password: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.CreateUser(dbPointer, db.CreateUserParams{
		Username:     username,
		Name:         username,
		TempPassword: hashedPassword,
	})

	if err != nil {
		log.Println("Error creating user: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userDetails := map[string]string{
		"username":     username,
		"tempPassword": tempPassword,
	}

	jsonData, err := json.Marshal(userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
