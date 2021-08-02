package db

import (
	"clients/cassandra"

	"github.com/ayush723/oauth-api_bookstore/src/domain/access_token"

	"github.com/ayush723/oauth-api_bookstore/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	//TODO: implement get access token from cassandra db
	_, err := cassandra.GetSession()
	if err != nil{
		panic(err)
	}
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
