package model

import "time"

// Auth ...
type Auth struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
}

// AuthWithJWT ...
type AuthWithJWT struct {
	Auth                *Auth  `json:"auth"`
	AuthToken           string `json:"auth_token"`
	AuthTokenExp        int64  `json:"auth_token_expired"`
	AuthTokenRefresh    string `json:"auth_token_refresh"`
	AuthTokenRefreshExp int64  `json:"auth_token_refresh_expired"`
}

// JWTToken ...
type JWTToken struct {
	Token   string
	Expired time.Time
}

// UserSession ...
type UserSession struct {
	ID       int64
	Username string
	Email    string
}
