package repository

import (
	"github.com/student/marketing-service-backend/services/common"
	"github.com/student/marketing-service-backend/services/project/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Project, error) {
	query := "SELECT * FROM project"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []internal.Project
	for rows.Next() {
		var project internal.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description,
			&project.ClientID, &project.ContractID); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (ir *Repository) Get(id int) (*internal.Project, error) {
	query := "SELECT * FROM project WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var project internal.Project
	if err := row.Scan(&project.ID, &project.Title, &project.Description,
		&project.ClientID, &project.ContractID); err != nil {
		return nil, err
	}

	return &project, nil
}

func (ir *Repository) Create(newProject *internal.Project) error {
	query := "INSERT INTO project (title, description, client_id, contract_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newProject.Title, newProject.Description,
		newProject.ClientID, newProject.ContractID).Scan(&newProject.ID)
	return err
}

func (ir *Repository) Update(id int, title string, description string, clientID int, contractID int) error {
	query := "UPDATE project SET title = $1, description = $2, contract_id = $3, client_id = $4 WHERE id = $5"
	_, err := ir.DB.Connection.Exec(query, title, description, contractID, clientID, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM project WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
