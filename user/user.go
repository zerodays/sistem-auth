package user

import (
	"github.com/zerodays/sistem-auth/permission"
	"github.com/zerodays/sistem-auth/token"
)

type User struct {
	UID         string
	Permissions []permission.Permission
}

// FromToken returns user from access token. If access token is not valid,
// error is returned.
func FromToken(accessToken string) (*User, error) {
	claims, err := token.Validate(accessToken)
	if err != nil {
		return nil, err
	}

	u := &User{
		UID:         claims.Subject,
		Permissions: claims.Permissions,
	}
	return u, nil
}

// HasPermission checks if user has permission perm.
func (u *User) HasPermission(perm permission.Permission) bool {
	for _, p := range u.Permissions {
		if p == perm {
			return true
		}
	}

	return false
}
