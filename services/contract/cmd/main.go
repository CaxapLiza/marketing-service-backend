package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/marketing-service-backend/services/contract/internal/handler"
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

	router.HandleFunc("/contracts", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/contracts/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/contracts", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/contracts/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/contracts/{id}", handler.Delete).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}
