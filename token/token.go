package token

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zerodays/sistem-auth/permission"
	"io/ioutil"
	"net/http"
)

var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	jwt.StandardClaims

	Permissions []permission.Type `json:"permissions"`
}

var signingKey *rsa.PublicKey

// Validate checks if token is valid and returns its claims.
func Validate(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
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
