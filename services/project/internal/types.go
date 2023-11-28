package internal

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ClientID    int    `json:"client_id"`
	ContractID  int    `json:"contract_id"`
}
