package access_token

import (
	"strings"

	"github.com/ayush723/oauth-api_bookstore/src/repository/rest"
	"github.com/ayush723/oauth-api_bookstore/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *errors.RestErr ){
	if err := request.Validate(); err != nil {
		return nil, err
	}
	//TODO: support both client-credentials and passwords

	// Authenticate the user against the users api:
	user, err := rest.NewRepository().LoginUser(request.Username,request.Password)
	if err != nil{
		return nil, err
	}
	//Generate new access token
	at := GetNewAccessToken(user.Id)
	at.Generate()
	//Save the new access token in Cassandra
	return &at,nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
