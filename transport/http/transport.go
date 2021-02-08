package http

import (
	"context"
	"net/http"

	"github.com/setiadijoe/go-auth-jwt-sample/endpoint"
	"github.com/setiadijoe/go-auth-jwt-sample/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type err interface {
	error() error
}

// MakeHandler ...
func MakeHandler(ctx context.Context, fs usecase.AuthI, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	processLogin := kithttp.NewServer(
		endpoint.MakeLogin(ctx, fs),
		decodeMakeLogin,
		encodeResponse,
		opts...,
	)

	processLogout := kithttp.NewServer(
		endpoint.MakeLogout(ctx, fs),
		decodeMakeLogout,
		encodeResponse,
		opts...,
	)

	processCheckJWT := kithttp.NewServer(
		endpoint.MakeCheckJWT(ctx, fs),
		decodeMakeCheckJWT,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/auth/login", processLogin).Methods("POST")
	r.Handle("/auth/logout", processLogout).Methods("POST")
	r.Handle("/auth", processCheckJWT).Methods("GET")

	return r

}
