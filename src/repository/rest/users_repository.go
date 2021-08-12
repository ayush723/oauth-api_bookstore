package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ayush723/oauth-api_bookstore/src/domain/users"
	"github.com/ayush723/utils-go_bookstore/rest_errors"

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
	LoginUser(string,string)(*users.User, rest_errors.RestErr)
}

type usersRepository struct{}

//NewRestUsersRepository returns a instance of type RestUsersRepositry interface(containing dbRepository struct) to mock the methods.
func NewRestUsersRepository() RestUsersRepository{
	return &usersRepository{}
}


//LoginUser consumes the api from users api(users/login)
func (r *usersRepository) LoginUser(email string, password string)(*users.User, rest_errors.RestErr){
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	
	response := usersRestClient.Post("/users/login",request)


	//checking error in response
	if response == nil || response.Request == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))

	}
	if response.StatusCode > 299{
		
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil{
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr 
	}
	// if response is ok, we unmarshal the response body
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil{
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", err)
	}
	return &user, nil

}