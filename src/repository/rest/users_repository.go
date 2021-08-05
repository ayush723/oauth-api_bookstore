package rest

import (
	"encoding/json"

	"github.com/ayush723/oauth-api_bookstore/src/domain/users"
	"github.com/ayush723/oauth-api_bookstore/src/utils/errors"
	"github.com/go-resty/resty/v2"
)

var (
	client = resty.New()
	
)
type RestUsersRepository interface{
	LoginUser(string,string)(*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository{
	return &usersRepository{}
}



func (r *usersRepository) LoginUser(email string, password string)(*users.User, *errors.RestErr){
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	
	response, err := client.R().
	SetHeader("Content-Type", "application/json").
	SetBody(request).
	Post("https://api.bookstore.com/users/login")

	// response, err := client.R().EnableTrace().Post("https://api.bookstore.com/users/login")
	if err != nil{
		return nil, errors.NewInternalServerError("error in the api")
	}



	if response == nil || response.Request == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")

	}
	if response.StatusCode() > 299{
		var restErr errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil{
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil{
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil

}