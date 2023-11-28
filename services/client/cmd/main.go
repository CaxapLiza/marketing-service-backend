package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/marketing-service-backend/services/client/internal/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	)

	router.Use(corsMiddleware)

	router.HandleFunc("/clients", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/clients/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/clients", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/clients/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/clients/{id}", handler.Delete).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8081", router))
}
