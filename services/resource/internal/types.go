package internal

type Resource struct {
	ID          int    `json:"id"`
	Link        string `json:"link"`
	Description string `json:"description"`
	ProjectID   int    `json:"project_id"`
}
