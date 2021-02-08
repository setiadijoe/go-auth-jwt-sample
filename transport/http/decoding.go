package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/setiadijoe/go-auth-jwt-sample/endpoint"
)

func decodeMakeLogin(ctx context.Context, r *http.Request) (interface{}, error) {
	request := &endpoint.UserLogin{}

	if err := json.NewDecoder(r.Body).Decode(request); nil != err {
		return nil, err
	}

	return request, nil
}

func decodeMakeLogout(ctx context.Context, r *http.Request) (interface{}, error) {
	request := &endpoint.UserLogout{}

	if err := json.NewDecoder(r.Body).Decode(request); nil != err {
		return nil, err
	}

	return request, nil
}

func decodeMakeCheckJWT(ctx context.Context, r *http.Request) (interface{}, error) {
	token := r.URL.Query().Get("token")

	return &endpoint.UserToken{
		Token: token,
	}, nil

}
