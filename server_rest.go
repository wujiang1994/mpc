package mpc

import (
	"log"
	"net"
	"net/http"
)

func (s *AppServer) serveREST() {
	defer s.wg.Done()
	conn, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Printf("listeners.Listen: %v", err)
		return
	}
	s.rest = &http.Server{
		Addr:              "0.0.0.0:5000",
		Handler:           s.AppGroup,
	}
	s.rest.Serve(conn)
}