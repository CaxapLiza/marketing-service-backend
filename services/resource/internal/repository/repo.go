package repository

import (
	"github.com/student/marketing-service-backend/services/common"
	"github.com/student/marketing-service-backend/services/resource/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.Resource, error) {
	query := "SELECT * FROM resource WHERE project_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []internal.Resource
	for rows.Next() {
		var resource internal.Resource
		if err := rows.Scan(&resource.ID, &resource.Link, &resource.Description, &resource.ProjectID); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}

func (ir *Repository) Get(id int) (*internal.Resource, error) {
	query := "SELECT * FROM resource WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var resource internal.Resource
	if err := row.Scan(&resource.ID, &resource.Link, &resource.Description, &resource.ProjectID); err != nil {
		return nil, err
	}

	return &resource, nil
}

func (ir *Repository) Create(newResource *internal.Resource) error {
	query := "INSERT INTO resource (link, description, project_id) VALUES ($1, $2, $3) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newResource.Link, newResource.Description, newResource.ProjectID).Scan(&newResource.ID)
	return err
}

func (ir *Repository) Update(id int, link string, description string, projectID int) error {
	query := "UPDATE resource SET link = $1, description = $2, project_id = $3 WHERE id = $4"
	_, err := ir.DB.Connection.Exec(query, link, description, projectID, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM resource WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
