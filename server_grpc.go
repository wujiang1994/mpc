package mpc

import (
	"google.golang.org/grpc"
	"net"
)

func (s *AppServer) serveGRPC() {
	defer s.wg.Done()

	grpcConfig := s.config.GRPCServerConfig()
	network, addr := grpcConfig.Bind()
	listener, err := net.Listen(network, addr)
	if err != nil {
		panic(err)
	}
	s.grpc = grpc.NewServer()
	s.mux.Lock()
	for _, grpcService := range s.grpcServices {
		grpcService.Register(s.grpc)
	}
	s.mux.Unlock()
	s.grpc.Serve(listener)
}
