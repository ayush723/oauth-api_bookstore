package access_token

import "utils/errors"

type Service interface{
	GetById(string)(*AccessToken, *errors.RestErr)
}

type service struct{
	
}