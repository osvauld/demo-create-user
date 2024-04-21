package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/osvauld/demo-user/db"
)

func createRandomUser() (string, string) {
	gofakeit.Seed(0)
	username := gofakeit.Username()
	tempPassword := gofakeit.Password(true, true, true, true, true, 12)

	return username, tempPassword
}

func createDemoUserHandler(w http.ResponseWriter, r *http.Request) {

	username, tempPassword := createRandomUser()

	_, err := db.CreateUser(dbPointer, db.CreateUserParams{
		Username:     username,
		Name:         username,
		TempPassword: tempPassword,
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
