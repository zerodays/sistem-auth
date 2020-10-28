package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/zerodays/sistem-auth/permission"
	"io/ioutil"
	"net/http"
)

type Claims struct {
	jwt.StandardClaims

	Permissions []permission.Permission `json:"permissions"`
}

var signingKey *rsa.PublicKey

// Validate checks if token is valid and returns its claims.
func Validate(accessToken string) (Claims, error) {
	token, err := jwt.Parse(accessToken, func(_ *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return Claims{}, err
	}

	return token.Claims.(Claims), nil
}

// LoadKey loads RSA public key used for verifying access tokens from specified url.
func LoadKey(url string) error {
	// Get key from specified url.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read server response.
	keyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse key.
	signingKey, err = jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	return err
}

// SetKey sets public key. When loading key from different url, use LoadKey.
func SetKey(key *rsa.PublicKey) {
	signingKey = key
}
