package model

import "time"

type Bank struct {
	ID        string    `json: "id"`
	Code      string    `json: "code"`
	Name      string    `json: "name"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func NewBank(code string, name string) (*Bank, error) {
	bank := new Bank{
		Code: code,
		Name: name,
	}

	return &bank, nil
}