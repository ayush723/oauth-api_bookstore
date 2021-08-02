package access_token

import "github.com/ayush723/oauth-api_bookstore/src/utils/errors"

type Repository interface{
	GetById(string)(*AccessToken, *errors.RestErr)	
}

type Service interface{
	GetById(string)(*AccessToken, *errors.RestErr)
}

type service struct{
	repository Repository
}

func NewService(repo Repository)Service{
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(string)(*AccessToken, *errors.RestErr){

}