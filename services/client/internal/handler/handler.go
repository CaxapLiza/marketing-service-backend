package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/student/marketing-service-backend/services/client/internal"
	"github.com/student/marketing-service-backend/services/client/internal/repository"
	"github.com/student/marketing-service-backend/services/common"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func GetList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	clients, err := repo.GetList()
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(clients)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Get(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	clients, err := repo.Get(id)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(clients)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var requestBody struct {
		Name                 string `json:"name"`
		TinOrKpp             string `json:"TIN_or_KPP"`
		Address              string `json:"address"`
		BIK                  string `json:"BIK"`
		CheckingAccount      string `json:"checking_account"`
		CorrespondentAccount string `json:"correspondent_account"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	client := &internal.Client{Name: requestBody.Name, TinOrKpp: requestBody.TinOrKpp,
		Address: requestBody.Address, BIK: requestBody.BIK,
		CheckingAccount: requestBody.CheckingAccount, CorrespondentAccount: requestBody.CorrespondentAccount}

	if err := repo.Create(client); err != nil {
		http.Error(w, "Error creating", http.StatusInternalServerError)
		log.Println("Create Error:", err)
		return
	}

	response, err := json.Marshal(client)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	var requestBody struct {
		Name                 string `json:"name"`
		TinOrKpp             string `json:"TIN_or_KPP"`
		Address              string `json:"address"`
		BIK                  string `json:"BIK"`
		CheckingAccount      string `json:"checking_account"`
		CorrespondentAccount string `json:"correspondent_account"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	if err := repo.Update(id, requestBody.Name, requestBody.TinOrKpp,
		requestBody.Address, requestBody.BIK,
		requestBody.CheckingAccount, requestBody.CorrespondentAccount); err != nil {
		http.Error(w, "Error updating", http.StatusInternalServerError)
		log.Println("Update Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	if err := repo.Delete(id); err != nil {
		http.Error(w, "Error deleting", http.StatusInternalServerError)
		log.Println("Delete Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
