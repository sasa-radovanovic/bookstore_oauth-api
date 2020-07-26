package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)
	assert.Equal(t, false, at.IsExpired(), "new access token should not be expired")
	assert.Equal(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.Equal(t, int64(0), at.UserID, "new access token should not have defined user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.Equal(t, true, at.IsExpired(), "access token created 3 hours from now should not be expired")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.Equal(t, false, at.IsExpired(), "access token created 3 hours from now should not be expired")
}

func TestExpirationTimeConstant(t *testing.T) {
	assert.Equal(t, 24, expirationTime, "Expiration time should be 24 hours")
}
