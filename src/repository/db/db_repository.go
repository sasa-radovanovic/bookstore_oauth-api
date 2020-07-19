package db

import (
	accesstoken "github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/access_token"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

// DatabaseRepository is a database repository
type DatabaseRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}
