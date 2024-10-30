package services

import (
	"encoding/json"
	"net/http"

	"github.com/Arnav2300/cal-set/internal/dtos"
	"github.com/Arnav2300/cal-set/internal/models"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		//generate jwt is all is good
		//return token

	}
}
