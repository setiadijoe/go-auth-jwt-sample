package endpoint

import (
	"context"

	"github.com/setiadijoe/go-auth-jwt-sample/usecase"

	"github.com/go-kit/kit/endpoint"
)

// MakeLogin ...
func MakeLogin(ctx context.Context, svc usecase.AuthI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqData := request.(*UserLogin)
		resp, err := svc.Login(ctx, reqData.Email, reqData.Password)
		if nil != err {
			return nil, err
		}

		return &Response{
			Data:    resp,
			Message: "user_succes_login",
		}, nil
	}
}

// MakeLogout ...
func MakeLogout(ctx context.Context, svc usecase.AuthI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqData := request.(*UserLogout)
		resp, err := svc.Logout(ctx, reqData.Email)
		if nil != err {
			return nil, err
		}

		return &Response{
			Data:    resp,
			Message: "user_success_logout",
		}, nil
	}
}

// MakeCheckJWT ...
func MakeCheckJWT(ctx context.Context, svc usecase.AuthI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqData := request.(*UserToken)
		resp, err := svc.CheckJWT(ctx, reqData.Token)
		if nil != err {
			return nil, err
		}

		return &Response{
			Data:    resp,
			Message: "user_session",
		}, nil
	}
}
