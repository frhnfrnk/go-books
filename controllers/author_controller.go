package controllers

import (
	"encoding/json"
	"github.com/frhnfrnk/go-books/models"
	"github.com/frhnfrnk/go-books/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func AttachAuthorRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/authors", GetAllAuthors(db)).Methods("GET")
	router.HandleFunc("/authors/{id}", GetAuthor(db)).Methods("GET")
	router.HandleFunc("/authors", CreateAuthor(db)).Methods("POST")
	router.HandleFunc("/authors/{id}", UpdateAuthor(db)).Methods("PUT")
	router.HandleFunc("/authors/{id}", DeleteAuthor(db)).Methods("DELETE")
}

func GetAllAuthors(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authors []models.Author

		// Query database to get all authors
		if err := db.Find(&authors).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the list of authors
		utils.RespondWithJSON(w, http.StatusOK, authors)
	}
}

func GetAuthor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		authorID := vars["id"]

		var author models.Author

		// Query database to get the author by ID
		if err := db.Where("id = ?", authorID).First(&author).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Author not found")
			return
		}

		// Respond with the author
		utils.RespondWithJSON(w, http.StatusOK, author)
	}
}

func CreateAuthor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author

		// Parse request body into the 'author' variable
		err := json.NewDecoder(r.Body).Decode(&author)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Create the author in the database
		if err := db.Create(&author).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the created author
		utils.RespondWithJSON(w, http.StatusCreated, author)
	}
}

func UpdateAuthor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		authorID := vars["id"]

		var updatedAuthor models.Author

		// Parse request body into the 'updatedAuthor' variable
		err := json.NewDecoder(r.Body).Decode(&updatedAuthor)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		var existingAuthor models.Author

		// Query database to get the existing author by ID
		if err := db.Where("id = ?", authorID).First(&existingAuthor).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Author not found")
			return
		}

		// Update the existing author in the database
		db.Model(&existingAuthor).Updates(updatedAuthor)

		// Respond with the updated author
		utils.RespondWithJSON(w, http.StatusOK, existingAuthor)
	}
}

func DeleteAuthor(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		authorID := vars["id"]

		var author models.Author

		// Query database to get the author by ID
		if err := db.Where("id = ?", authorID).First(&author).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Author not found")
			return
		}

		// Delete the author from the database
		db.Delete(&author)

		// Respond with a success message
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Author deleted successfully"})
	}
}
