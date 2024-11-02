package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Arnav2300/cal-set/api/db"
	"github.com/Arnav2300/cal-set/api/dtos"
	"github.com/Arnav2300/cal-set/api/models"
	"github.com/Arnav2300/cal-set/api/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {

	db := db.GetDatabase()
	var input dtos.LoginInput

	//parse JSON input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//find user in DB
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	//check if password is correct
	if !utils.CheckPassword(input.Password, *user.Password) {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	//generate jwt if all is good
	token, err := utils.GenerateToken(user.Email, user.Role)
	if err != nil {
		log.Printf("Error while generating token: %s", err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	//return token
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
