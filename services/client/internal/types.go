package internal

type Client struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	TinOrKpp             string `json:"TIN_or_KPP"`
	Address              string `json:"address"`
	BIK                  string `json:"BIK"`
	CheckingAccount      string `json:"checking_account"`
	CorrespondentAccount string `json:"correspondent_account"`
}
