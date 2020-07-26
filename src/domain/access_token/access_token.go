package accesstoken

import (
	"fmt"
	"strings"
	"time"

	cryptoutils "github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

// AccessToken domain object
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// AtRequest domain object
type AtRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// Used for grant_type password
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for grant_type client_credentials
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// GetNewAccessToken retrieves new access token
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired checks expiration
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}

// Validate Access Token Request
func (atRequest *AtRequest) Validate() *errors.RestErr {
	switch atRequest.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	// TODO
	return nil
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

// Generate generates new access token
func (at *AccessToken) Generate() {
	at.AccessToken = cryptoutils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
