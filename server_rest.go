package mpc

import (
	"net"
	"net/http"
	"time"
)

func (s *AppServer) serveREST() {
	defer s.wg.Done()
	restConfig := s.config.RestServerConfig()
	network, addr := s.config.RestServerConfig().Bind()
	conn, err := net.Listen(network, addr)
	if err != nil {
		panic(err)
	}
	s.rest = &http.Server{
		Addr:              addr,
		Handler:           s.AppGroup,
		ReadTimeout:       time.Duration(restConfig.RequestTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(restConfig.RequestTimeout) * time.Second,
		WriteTimeout:      time.Duration(restConfig.ResponseTimeout) * time.Second,
		IdleTimeout:       time.Duration(restConfig.ConnectTimeout) * time.Second,
		MaxHeaderBytes:    restConfig.MaxHeaderBytes,
	}
	if s.config.RestServerConfig().Ssl {
		s.rest.ServeTLS(conn, restConfig.SslCert, restConfig.SslKey)
	} else {
		s.rest.Serve(conn)
	}
}
