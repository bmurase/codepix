package usecase

import (
	"errors"

	"github.com/bmurase/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountID string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)

	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.Register(pixKey)

	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
