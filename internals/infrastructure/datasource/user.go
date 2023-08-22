package datasource

import (
	"github.com/gushikem01/go-handson/internals/config"
	"github.com/gushikem01/go-handson/internals/domain/repository"
)

type userRepository struct {
	pg *config.PostgresClient
	tx config.Transaction
}

func NewUserRepository(pg *config.PostgresClient, tx config.Transaction) repository.UserRepository {
	return &userRepository{
		pg: pg,
		tx: tx,
	}
}
