package accesstoken

import (
	"strings"
	"time"

	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken domain object
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// GetNewAccessToken retrieves new access token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired checks expiration
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}

// Validate validates the data
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}
