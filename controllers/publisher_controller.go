package controllers

import (
	"encoding/json"
	"github.com/frhnfrnk/go-books/models"
	"github.com/frhnfrnk/go-books/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func AttachPublisherRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/publishers", GetAllPublishers(db)).Methods("GET")
	router.HandleFunc("/publishers/{id}", GetPublisher(db)).Methods("GET")
	router.HandleFunc("/publishers", CreatePublisher(db)).Methods("POST")
	router.HandleFunc("/publishers/{id}", UpdatePublisher(db)).Methods("PUT")
	router.HandleFunc("/publishers/{id}", DeletePublisher(db)).Methods("DELETE")
}

func GetAllPublishers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var publishers []models.Publisher

		// Query database to get all publishers
		if err := db.Find(&publishers).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the list of publishers
		utils.RespondWithJSON(w, http.StatusOK, publishers)
	}
}

func GetPublisher(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		publisherID := vars["id"]

		var publisher models.Publisher

		// Query database to get the publisher by ID
		if err := db.Where("id = ?", publisherID).First(&publisher).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Publisher not found")
			return
		}

		// Respond with the publisher
		utils.RespondWithJSON(w, http.StatusOK, publisher)
	}
}

func CreatePublisher(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var publisher models.Publisher

		// Parse request body into the 'publisher' variable
		err := json.NewDecoder(r.Body).Decode(&publisher)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Create the publisher in the database
		if err := db.Create(&publisher).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the created publisher
		utils.RespondWithJSON(w, http.StatusCreated, publisher)
	}
}

func UpdatePublisher(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		publisherID := vars["id"]

		var updatedPublisher models.Publisher

		// Parse request body into the 'updatedPublisher' variable
		err := json.NewDecoder(r.Body).Decode(&updatedPublisher)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		var existingPublisher models.Publisher

		// Query database to get the existing publisher by ID
		if err := db.Where("id = ?", publisherID).First(&existingPublisher).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Publisher not found")
			return
		}

		// Update the existing publisher in the database
		db.Model(&existingPublisher).Updates(updatedPublisher)

		// Respond with the updated publisher
		utils.RespondWithJSON(w, http.StatusOK, existingPublisher)
	}
}

func DeletePublisher(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		publisherID := vars["id"]

		var publisher models.Publisher

		// Query database to get the publisher by ID
		if err := db.Where("id = ?", publisherID).First(&publisher).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Publisher not found")
			return
		}

		// Delete the publisher from the database
		db.Delete(&publisher)

		// Respond with a success message
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Publisher deleted successfully"})
	}
}
