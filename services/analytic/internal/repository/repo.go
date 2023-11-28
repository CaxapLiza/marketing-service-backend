package repository

import (
	"github.com/student/marketing-service-backend/services/analytic/internal"
	"github.com/student/marketing-service-backend/services/common"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.Analytic, error) {
	query := "SELECT * FROM analytic WHERE project_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var analytics []internal.Analytic
	for rows.Next() {
		var analytic internal.Analytic
		if err := rows.Scan(&analytic.ID, &analytic.CTR, &analytic.ConversationRate,
			&analytic.ROI, &analytic.CAC, &analytic.LTV, &analytic.ProjectID); err != nil {
			return nil, err
		}
		analytics = append(analytics, analytic)
	}

	return analytics, nil
}

func (ir *Repository) Get(id int) (*internal.Analytic, error) {
	query := "SELECT * FROM analytic WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var analytic internal.Analytic
	if err := row.Scan(&analytic.ID, &analytic.CTR, &analytic.ConversationRate, &analytic.ROI,
		&analytic.CAC, &analytic.LTV, &analytic.ProjectID); err != nil {
		return nil, err
	}

	return &analytic, nil
}

func (ir *Repository) Create(newAnalytic *internal.Analytic) error {
	query := "INSERT INTO analytic (ctr, conversation_rate, roi, cac, ltv, project_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newAnalytic.CTR, newAnalytic.ConversationRate, newAnalytic.ROI,
		newAnalytic.CAC, newAnalytic.LTV, newAnalytic.ProjectID).Scan(&newAnalytic.ID)
	return err
}

func (ir *Repository) Update(id int, ctr float32, conversationRate float32, roi float32, cac float32, ltv float32) error {
	query := "UPDATE analytic SET ctr = $1, conversation_rate = $2, roi = $3, cac = $4, ltv = $5 WHERE id = $6"
	_, err := ir.DB.Connection.Exec(query, ctr, conversationRate, roi, cac, ltv, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM analytic WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
