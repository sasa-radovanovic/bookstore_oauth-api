package db

import (
	"github.com/gocql/gocql"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/clients/cassandra"
	accesstoken "github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/access_token"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// NewRepository returns new database repository
func NewRepository() DatabaseRepository {
	return &dbRepository{}
}

// DatabaseRepository is a database repository
type DatabaseRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	var result accesstoken.AccessToken
	if dbErr := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); dbErr != nil {
		if dbErr == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token with given id found")
		}
		return nil, errors.NewInternalServerError(dbErr.Error())
	}

	return &result, nil
}

func (r *dbRepository) UpdateExpirationTime(accessToken accesstoken.AccessToken) *errors.RestErr {
	if dbErr := cassandra.GetSession().Query(queryUpdateExpires,
		accessToken.Expires,
		accessToken.AccessToken,
	).Exec(); dbErr != nil {
		return errors.NewInternalServerError(dbErr.Error())
	}

	return nil
}

func (r *dbRepository) Create(accessToken accesstoken.AccessToken) *errors.RestErr {
	if dbErr := cassandra.GetSession().Query(queryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserID,
		accessToken.ClientID,
		accessToken.Expires,
	).Exec(); dbErr != nil {
		return errors.NewInternalServerError(dbErr.Error())
	}

	return nil
}
