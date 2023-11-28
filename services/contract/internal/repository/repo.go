package repository

import (
	"github.com/student/marketing-service-backend/services/common"
	"github.com/student/marketing-service-backend/services/contract/internal"
	"time"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Contract, error) {
	query := "SELECT * FROM contract"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []internal.Contract
	for rows.Next() {
		var contract internal.Contract
		if err := rows.Scan(&contract.ID, &contract.DateConcluded, &contract.StartDate, &contract.EndDate,
			&contract.Price, &contract.Text, &contract.ClientID); err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (ir *Repository) Get(id int) (*internal.Contract, error) {
	query := "SELECT * FROM contract WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var contract internal.Contract
	if err := row.Scan(&contract.ID, &contract.DateConcluded, &contract.StartDate, &contract.EndDate,
		&contract.Price, &contract.Text, &contract.ClientID); err != nil {
		return nil, err
	}

	return &contract, nil
}

func (ir *Repository) Create(newContract *internal.Contract) error {
	query := "INSERT INTO contract (date_concluded, start_date, end_date, price, text, client_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newContract.DateConcluded, newContract.StartDate, newContract.EndDate,
		newContract.Price, newContract.Text, newContract.ClientID).Scan(&newContract.ID)
	return err
}

func (ir *Repository) Update(id int, dateConcluded time.Time, startDate time.Time, endDate time.Time,
	price float32, text string, clientId int) error {
	query := "UPDATE contract SET date_concluded = $1, start_date = $2, end_date = $3, price = $4, text = $5, client_id = $6 WHERE id = $7"
	_, err := ir.DB.Connection.Exec(query, dateConcluded, startDate, endDate, price, text, clientId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM contract WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
