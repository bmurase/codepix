package repository

import (
	"fmt"

	"github.com/bmurase/codepix/domain/model"
	gorm "gorm.io/gorm"
)

type TransactionRepository struct {
	Db *gorm.DB
}

func (r *TransactionRepository) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found with given id")
	}

	return &transaction, nil
}
