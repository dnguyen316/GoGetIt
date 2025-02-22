package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/dnguyen316/GoGetIt/server/internal/generated/grpc/go_get_it"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
}

func NewServer() Server {
	return &server{}
}

func (s *server) Start(ctx context.Context) error {
	mux := runtime.NewServeMux()
	if err := go_get_it.RegisterGoGetItServiceHandlerFromEndpoint(
		ctx,
		mux,
		"0.0.0.0:8080",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}); err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}
