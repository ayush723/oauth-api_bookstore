package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases..")
	// client := &http.Client{Transport: &http.Transport{TLSHandshakeTimeout: 60 * time.Second}}
	// httpmock.ActivateNonDefault(client)
	os.Exit(m.Run())
}

func TestLoginuserTimeoutFromApi(t *testing.T) {

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {}

func TestLoginUserNoError(t *testing.T) {}
