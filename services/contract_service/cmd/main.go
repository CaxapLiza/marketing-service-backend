package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/marketing-service-backend/services/contract_service/internal/handler"
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

	router.HandleFunc("/contract_services/list/{id}", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/contract_services/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/contract_services", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/contract_services/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/contract_services/{id}", handler.Delete).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8083", router))
}
