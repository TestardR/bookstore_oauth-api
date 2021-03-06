package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, false, at.IsExpired())
	assert.EqualValues(t, "", at.AccessToken)
	assert.True(t, 0 == at.UserId)
}

func TestGetNewAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired())
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired())
}
