package db

import (
	"errors"

	"github.com/ayush723/oauth-api_bookstore/src/clients/cassandra"
	"github.com/ayush723/oauth-api_bookstore/src/domain/access_token"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

//NewRepository returns a instance of type DbRepositry interface(containing dbRepository struct) to mock the methods.
func NewRepository() DbRepository {
	return &dbRepository{}
}

//DbRepository implements all methods on access_token on database level
type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestErr)
	Create(access_token.AccessToken) rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestErr
}

type dbRepository struct {
}

//GetById get access_token details from db
func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestErr) {

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no acess token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error when trying to get current id", err)
	}
	return &result, nil
}

//Create creates new access_token in database
func (r *dbRepository) Create(at access_token.AccessToken) rest_errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	return nil
}

//UpdateExpirationTime updates the expiration time of gien access_token
func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) rest_errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return rest_errors.NewInternalServerError("error when trying to update current resource", errors.New("database errors"))
	}
	return nil
}
