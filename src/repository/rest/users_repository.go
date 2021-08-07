package rest

import (
	"encoding/json"
	"time"

	"github.com/ayush723/oauth-api_bookstore/src/domain/users"
	"github.com/ayush723/oauth-api_bookstore/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
	
)

//RestusersRepository implements the Loginuser method
type RestUsersRepository interface{
	LoginUser(string,string)(*users.User, *errors.RestErr)
}

type usersRepository struct{}

//NewRestUsersRepository returns a instance of type RestUsersRepositry interface(containing dbRepository struct) to mock the methods.
func NewRestUsersRepository() RestUsersRepository{
	return &usersRepository{}
}


//LoginUser consumes the api from users api(users/login)
func (r *usersRepository) LoginUser(email string, password string)(*users.User, *errors.RestErr){
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	
	response := usersRestClient.Post("/users/login",request)


	//checking error in response
	if response == nil || response.Request == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")

	}
	if response.StatusCode > 299{
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil{
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	// if response is ok, we unmarshal the response body
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil{
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil

}