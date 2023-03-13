package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jeffreylean/alir/config"
	alir "github.com/jeffreylean/alir/proto"
	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	grpcPort int32
	cancel   context.CancelFunc
}

func Start(c *config.Config) *Server {
	server := &Server{
		server:   grpc.NewServer(),
		grpcPort: c.Port,
	}
	alir.RegisterIngestServiceServer(server.server, server)

	return server
}

func (s *Server) Ingest(ctx context.Context, request *alir.IngestRequest) (*alir.IngestResponse, error) {
	return nil, nil
}

func (s *Server) Listen(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		return err
	}

	log.Printf("server: listening for grpc on :%d", s.grpcPort)
	if err := s.server.Serve(listen); err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()
	s.cancel()
}
