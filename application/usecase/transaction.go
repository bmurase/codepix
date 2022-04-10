package usecase

import (
	"errors"

	"github.com/bmurase/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixKeyRepository      model.PixKeyRepositoryInterface
}

func (t *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := t.PixKeyRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixKeyRepository.FindByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Register(transaction)

	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

func (t *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Confirm()

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionID string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Complete()

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Cancel(transactionID string, description string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionID)

	if err != nil {
		return nil, err
	}

	transaction.Cancel(description)

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
