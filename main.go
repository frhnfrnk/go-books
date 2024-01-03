package main

import (
	"fmt"
	"github.com/frhnfrnk/go-books/controllers"
	"github.com/frhnfrnk/go-books/database"
	"github.com/frhnfrnk/go-books/middlewares"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Initialize database
	db := database.InitDB()
	defer db.Close()

	// Apply migrations
	database.RunMigrations(db)

	router.Use(middlewares.AuthenticationMiddleware)

	// Attach routes
	controllers.AttachBookRoutes(router, db)
	controllers.AttachAuthorRoutes(router, db)
	controllers.AttachPublisherRoutes(router, db)
	controllers.AttachUserRoute(router, db)
	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
