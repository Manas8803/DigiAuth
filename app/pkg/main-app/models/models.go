package models

type Wallet struct {
	Id    int    `json:id"`
	Did   string `json:"did"`
	Email string `json:"email" validate:"required"`
}

type Certificate struct {
	Id           int    `json:id"`
	WalletId     string `json:"wallet"`
	CredentialId string `json:"credential_id"`
	IssueDate    string `json:"issue_date"`
}
