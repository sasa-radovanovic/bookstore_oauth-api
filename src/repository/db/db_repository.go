package db

import (
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/clients/cassandra"
	accesstoken "github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/access_token"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

// NewRepository returns new database repository
func NewRepository() DatabaseRepository {
	return &dbRepository{}
}

// DatabaseRepository is a database repository
type DatabaseRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	// TODO implement get access token from Cassandra DB
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
