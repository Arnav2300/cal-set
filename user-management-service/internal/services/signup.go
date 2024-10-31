package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Arnav2300/cal-set/internal/db"
	"github.com/Arnav2300/cal-set/internal/dtos"
	"github.com/Arnav2300/cal-set/internal/models"
	"github.com/Arnav2300/cal-set/internal/utils"
)

func Singup(w http.ResponseWriter, r *http.Request) {
	db := db.GetDatabase()
	var input dtos.SignupInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	var preExistingUser models.User
	db.Where("email = ?", input.Email).First(&preExistingUser)

	if preExistingUser.Email != "" {
		http.Error(w, "Email already in use!", http.StatusOK)
		return
	}

	hashedPass, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  &hashedPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully!"})

}
