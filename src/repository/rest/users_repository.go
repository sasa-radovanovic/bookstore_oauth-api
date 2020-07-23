package rest

import (
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/users"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
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

	return nil, nil
}
