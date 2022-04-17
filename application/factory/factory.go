package factory

import (
	"github.com/bmurase/codepix/application/usecase"
	"github.com/bmurase/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepository{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixKeyRepository:      pixRepository,
	}

	return transactionUseCase
}
