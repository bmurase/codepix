package model

type Account struct {
	Base      `valid:"required"`
	OwnerName string `json: "owner_name" valid:"notnull"`
	Bank      *Bank  `valid: "-"`
	Number    string `json:"number" valid:"notnull"`
}
