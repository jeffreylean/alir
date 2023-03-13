package server

import (
	"context"

	"github.com/jeffreylean/alir/config"
	alir "github.com/jeffreylean/alir/proto"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

func Start(c *config.Config) *Server {
	server := &Server{
		server: grpc.NewServer(),
	}
	alir.RegisterIngestServiceServer(server.server, server)

	return server
}
func (s *Server) Ingest(ctx context.Context, request *alir.IngestRequest) (*alir.IngestResponse, error) {
	return nil, nil
}
