package db

import (
	"github.com/TestardR/bookstore_oauth-api/src/domain/access_token"
	"github.com/TestardR/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func New() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
