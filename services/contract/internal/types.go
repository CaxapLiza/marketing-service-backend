package internal

import "time"

type Contract struct {
	ID            int       `json:"id"`
	DateConcluded time.Time `json:"date_concluded"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float32   `json:"price"`
	Text          string    `json:"text"`
	ClientID      int       `json:"client_id"`
}
