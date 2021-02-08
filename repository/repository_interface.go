package repository

import (
	"context"

	"github.com/setiadijoe/go-auth-jwt-sample/model"
)

// AuthorizationI ...
type AuthorizationI interface {
	Login(ctx context.Context, email string, password string) (*model.Auth, error)
	Logout(ctx context.Context, email string) (*model.Auth, error)
}
