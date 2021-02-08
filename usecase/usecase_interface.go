package usecase

import (
	"context"

	"github.com/setiadijoe/go-auth-jwt-sample/model"
)

// AuthI ...
type AuthI interface {
	Login(ctx context.Context, email string, password string) (*model.AuthWithJWT, error)
	CreateJWT(ctx context.Context, duration int, dataPayload map[string]interface{}, signingKey []byte) (*model.JWTToken, error)
	Logout(ctx context.Context, email string) (*model.AuthWithJWT, error)
	CheckJWT(ctx context.Context, token string) (*model.UserSession, error)
}
