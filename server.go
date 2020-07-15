package mpc

import (
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"sync"
)

type AppServer struct {
	*AppGroup

	config    Configer
	group     Grouper
	requestID string

	mux  sync.RWMutex
	wg   sync.WaitGroup
	sig  chan os.Signal
	once sync.Once
	rest *http.Server
	grpc *grpc.Server

	restListener net.Listener
	grpcServices []GRPCService
}

func NewAppServer(runMode, cfgPath string) *AppServer {
	config, err := NewAppConfig(runMode, cfgPath)
	if err != nil {
		panic(err)
	}
	SetupLogger(config.Logger)

	srv := &AppServer{
		config:       config,
		requestID:    "default",
		grpcServices: make([]GRPCService, 0),
	}

	srv.once.Do(func() {
		srv.AppGroup = NewAppGroup("/", srv)
		srv.group = &AppGroup{
			server: srv,
			prefix: "/",
		}
	})
	return srv
}

func (s *AppServer) Mode() RunMode {
	return s.config.RunMode()
}

func (s *AppServer) Config() Configer {
	return s.config
}

func (s *AppServer) NewServer(svc RestService) *AppServer {
	svc.Init(s.config, s.group)
	svc.Filters()
	svc.Resources()

	return s
}

func (s *AppServer) Run() {
	s.wg.Add(2)
	go s.serveREST()
	go s.serveGRPC()
	s.wg.Wait()
}
