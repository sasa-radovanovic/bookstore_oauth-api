package accesstoken

import (
	"strings"

	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

// Service is a service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService creates a new instance of service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetById retrieves by ID
func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
