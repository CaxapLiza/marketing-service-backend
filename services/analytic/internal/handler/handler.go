package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/student/marketing-service-backend/services/analytic/internal"
	"github.com/student/marketing-service-backend/services/analytic/internal/repository"
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

	analytics, err := repo.GetList(id)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(analytics)
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

	analytics, err := repo.Get(id)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(analytics)
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
		NumberOfImpressions             int     `json:"number_of_impressions"`
		NumberOfClicks                  int     `json:"number_of_clicks"`
		NumberOfVisitors                int     `json:"number_of_visitors"`
		NumberOfConversions             int     `json:"number_of_conversions"`
		Profit                          float32 `json:"profit"`
		Costs                           float32 `json:"costs"`
		NumberOfNewCustomers            int     `json:"number_of_new_customers"`
		MarketingCosts                  int     `json:"marketing_costs"`
		AverageAnnualRevenuePerCustomer float32 `json:"average_annual_revenue_per_customer"`
		AverageCustomerLifespan         float32 `json:"average_customer_lifespan"`
		ProjectID                       int     `json:"project_id"`
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

	analytic := &internal.Analytic{
		CTR:              float32(requestBody.NumberOfClicks / (requestBody.NumberOfImpressions * 100)),
		ConversationRate: float32(requestBody.NumberOfConversions / (requestBody.NumberOfVisitors * 100)),
		ROI:              (requestBody.Profit - requestBody.Costs) / (requestBody.Costs * 100),
		CAC:              float32(requestBody.MarketingCosts / requestBody.NumberOfNewCustomers),
		LTV:              requestBody.AverageCustomerLifespan * requestBody.AverageAnnualRevenuePerCustomer,
		ProjectID:        requestBody.ProjectID,
	}

	if err := repo.Create(analytic); err != nil {
		http.Error(w, "Error creating", http.StatusInternalServerError)
		log.Println("Create Error:", err)
		return
	}

	response, err := json.Marshal(analytic)
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
		CTR              float32 `json:"ctr"`
		ConversationRate float32 `json:"conversation_rate"`
		ROI              float32 `json:"roi"`
		CAC              float32 `json:"cac"`
		LTV              float32 `json:"ltv"`
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

	if err := repo.Update(id, requestBody.CTR, requestBody.ConversationRate,
		requestBody.ROI, requestBody.CAC, requestBody.LTV); err != nil {
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
