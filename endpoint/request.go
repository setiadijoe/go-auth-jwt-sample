package endpoint

// UserLogin ...
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLogout ...
type UserLogout struct {
	Email string `json:"email"`
}

// UserToken ...
type UserToken struct {
	Token string `json:"token"`
}
