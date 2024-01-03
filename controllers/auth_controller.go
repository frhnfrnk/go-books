package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/frhnfrnk/go-books/middlewares"
	"github.com/frhnfrnk/go-books/models"
	"github.com/frhnfrnk/go-books/utils"
	"github.com/jinzhu/gorm"
)

func AttachUserRoute(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/login", Login(db)).Methods("POST")
	router.HandleFunc("/register", Register(db)).Methods("POST")
}

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := user.Validate(); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		token, err := middlewares.GenerateToken(user.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the generated token
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
	}
}

func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if err := user.Validate(); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		if isUsernameTaken(db, user.Username) {
			utils.RespondWithError(w, http.StatusConflict, "Username is already taken")
			return
		}

		if err := db.Create(&user).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		token, err := middlewares.GenerateToken(user.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"token": token})
	}
}

// isUsernameTaken checks if the username is already taken
func isUsernameTaken(db *gorm.DB, username string) bool {
	var existingUser models.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err != nil {
		return false
	}
	return true
}
