package repository

import (
	"github.com/student/marketing-service-backend/services/common"
	"github.com/student/marketing-service-backend/services/contract_service/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(contractID int) ([]internal.ContractService, error) {
	query := "SELECT * FROM contract_service WHERE contract_id = $1"
	rows, err := ir.DB.Connection.Query(query, contractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contractServices []internal.ContractService
	for rows.Next() {
		var contractService internal.ContractService
		if err := rows.Scan(&contractService.ID, &contractService.ContractID, &contractService.ServiceID); err != nil {
			return nil, err
		}
		contractServices = append(contractServices, contractService)
	}

	return contractServices, nil
}

func (ir *Repository) Get(id int) (*internal.ContractService, error) {
	query := "SELECT * FROM contract_service WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var contractService internal.ContractService
	if err := row.Scan(&contractService.ID, &contractService.ContractID, &contractService.ServiceID); err != nil {
		return nil, err
	}

	return &contractService, nil
}

func (ir *Repository) Create(newContractService *internal.ContractService) error {
	query := "INSERT INTO contract_service (contract_id, service_id) VALUES ($1, $2) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newContractService.ContractID, newContractService.ServiceID).Scan(&newContractService.ID)
	return err
}

func (ir *Repository) Update(id int, contractID int, serviceID int) error {
	query := "UPDATE contract_service SET contract_id = $1, service_id = $2 WHERE id = $3"
	_, err := ir.DB.Connection.Exec(query, contractID, serviceID, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM contract_service WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
