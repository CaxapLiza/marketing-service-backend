package repository

import (
	"github.com/student/marketing-service-backend/services/client/internal"
	"github.com/student/marketing-service-backend/services/common"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Client, error) {
	query := "SELECT * FROM client"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []internal.Client
	for rows.Next() {
		var client internal.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.TinOrKpp, &client.Address,
			&client.BIK, &client.CheckingAccount, &client.CorrespondentAccount); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (ir *Repository) Get(id int) (*internal.Client, error) {
	query := "SELECT * FROM client WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var client internal.Client
	if err := row.Scan(&client.ID, &client.Name); err != nil {
		return nil, err
	}

	return &client, nil
}

func (ir *Repository) Create(newClient *internal.Client) error {
	query := "INSERT INTO client (name, tin_or_kpp, address, bik, checking_account, correspondent_account) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newClient.Name, newClient.TinOrKpp, newClient.Address,
		newClient.BIK, newClient.CheckingAccount, newClient.CorrespondentAccount).Scan(&newClient.ID)
	return err
}

func (ir *Repository) Update(id int, name string, tinOrKpp string, address string,
	bik string, checkAcc string, corrAcc string) error {
	query := "UPDATE client SET name = $1, tin_or_kpp = $2, address = $3, bik = $4, checking_account = $5, correspondent_account = $6 WHERE id = $7"
	_, err := ir.DB.Connection.Exec(query, name, tinOrKpp, address, bik, checkAcc, corrAcc, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM client WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
