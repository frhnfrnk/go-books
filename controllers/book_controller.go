package controllers

import (
	"encoding/json"
	"github.com/frhnfrnk/go-books/models"
	"github.com/frhnfrnk/go-books/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func AttachBookRoutes(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/books", GetAllBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook(db)).Methods("GET")
	router.HandleFunc("/books", CreateBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook(db)).Methods("DELETE")
}

func GetAllBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []models.Book

		// Query database to get all books
		if err := db.Find(&books).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the list of books
		utils.RespondWithJSON(w, http.StatusOK, books)
	}
}

func GetBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]

		var book models.Book

		// Query database to get the book by ID
		if err := db.Where("id = ?", bookID).First(&book).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Book not found")
			return
		}

		// Respond with the book
		utils.RespondWithJSON(w, http.StatusOK, book)
	}
}

func CreateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		// Parse request body into the 'book' variable
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Validate the book model
		if err := book.Validate(); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Create the book in the database
		if err := db.Create(&book).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Update the publisher with the new book
		var publisher models.Publisher
		if err := db.First(&publisher, book.PublisherID).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		publisher.Books = append(publisher.Books, book)
		if err := db.Save(&publisher).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Update the author with the new book
		var author models.Author
		if err := db.First(&author, book.AuthorID).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		author.Books = append(author.Books, book)
		if err := db.Save(&author).Error; err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Respond with the created book
		utils.RespondWithJSON(w, http.StatusCreated, book)
	}
}

func UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]

		var updatedBook models.Book

		// Parse request body into the 'updatedBook' variable
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		var existingBook models.Book

		// Query database to get the existing book by ID
		if err := db.Where("id = ?", bookID).First(&existingBook).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Book not found")
			return
		}

		// Update the existing book in the database
		db.Model(&existingBook).Updates(updatedBook)

		// Respond with the updated book
		utils.RespondWithJSON(w, http.StatusOK, existingBook)
	}
}

func DeleteBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]

		var book models.Book

		// Query database to get the book by ID
		if err := db.Where("id = ?", bookID).First(&book).Error; err != nil {
			utils.RespondWithError(w, http.StatusNotFound, "Book not found")
			return
		}

		// Delete the book from the database
		db.Delete(&book)

		// Respond with a success message
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
	}
}
