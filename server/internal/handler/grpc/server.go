package grpc

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"google.golang.org/grpc"

	"github.com/dnguyen316/GoGetIt/server/internal/generated/grpc/go_get_it"
)

type Server interface {
	Start(ctx context.Context) error
}

type server struct {
	handler go_get_it.GoGetItServiceServer
}

func NewServer(
	handler go_get_it.GoGetItServiceServer,
) Server {
	return &server{
		handler: handler,
	}
}

func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return err
	}

	defer listener.Close()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			validator.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			validator.StreamServerInterceptor(),
		),
	)
	go_get_it.RegisterGoGetItServiceServer(server, s.handler)
	return server.Serve(listener)

}
