package accesstoken

import (
	"strings"

	"github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/users"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

// RestRepository interface
type RestRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

// Service is a service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AtRequest) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository     Repository
	restRepository RestRepository
}

// NewService creates a new instance of service
func NewService(repo Repository, restRepo RestRepository) Service {
	return &service{
		repository:     repo,
		restRepository: restRepo,
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

func (s *service) Create(atRequest AtRequest) (*AccessToken, *errors.RestErr) {

	if err := atRequest.Validate(); err != nil {
		return nil, err
	}
	// TODO Support client_credentials and password
	user, err := s.restRepository.LoginUser(atRequest.Username, atRequest.Password)

	if err != nil {
		return nil, err
	}
	at := GetNewAccessToken(user.ID)
	at.Generate()

	if err := s.repository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
