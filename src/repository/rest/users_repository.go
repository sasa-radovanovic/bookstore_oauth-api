package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-restclient/rest"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/users"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

// UsersRepository is the user repository
type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRestRepository struct{}

// NewRepository returns new database repository
func NewRepository() UsersRepository {
	return &usersRestRepository{}
}

func (r *usersRestRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		fmt.Println("Is error != nil?", err != nil)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error trying to unmarshal users response")
	}
	return &user, nil
}
