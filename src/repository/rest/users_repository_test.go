package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func init() {

}

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Something", "code": "404", "status": "404"}`,
	})
	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials - 1", "error": "invalid user credentials received", "code": 404}`,
	})
	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code)
	assert.EqualValues(t, "invalid login credentials - 1", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"sssss", "first_name": 123, "last_name": "Radovanovic", "email": "sasa@sasa.rs"}`,
	})
	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	fmt.Println("Error data", err.Code, err.Message, err.Error)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, "error trying to unmarshal users response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":123, "first_name": "Sasa", "last_name": "Radovanovic", "email": "sasa@sasa.rs"}`,
	})
	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "Sasa", user.FirstName)
	assert.EqualValues(t, "Radovanovic", user.LastName)
	assert.EqualValues(t, "sasa@sasa.rs", user.Email)
	assert.EqualValues(t, 123, user.ID)
}
