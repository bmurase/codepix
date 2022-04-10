package model

import "github.com/asaskevich/govalidator"

type Account struct {
	Base      `valid:"required"`
	OwnerName string `json: "owner_name" valid:"notnull"`
	Bank      *Bank  `valid: "-"`
	Number    string `json:"number" valid:"notnull"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}
