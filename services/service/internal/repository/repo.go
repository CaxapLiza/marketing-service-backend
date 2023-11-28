package repository

import (
	"github.com/student/marketing-service-backend/services/common"
	"github.com/student/marketing-service-backend/services/service/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Service, error) {
	query := "SELECT * FROM service"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []internal.Service
	for rows.Next() {
		var service internal.Service
		if err := rows.Scan(&service.ID, &service.Description); err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

func (ir *Repository) Get(id int) (*internal.Service, error) {
	query := "SELECT * FROM service WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var service internal.Service
	if err := row.Scan(&service.ID, &service.Description); err != nil {
		return nil, err
	}

	return &service, nil
}

func (ir *Repository) Create(newService *internal.Service) error {
	query := "INSERT INTO service (description) VALUES ($1) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newService.Description).Scan(&newService.ID)
	return err
}

func (ir *Repository) Update(id int, description string) error {
	query := "UPDATE service SET description = $1 WHERE id = $2"
	_, err := ir.DB.Connection.Exec(query, description, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM service WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
