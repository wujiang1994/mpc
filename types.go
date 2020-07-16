package mpc

import (
	"google.golang.org/grpc"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Configer interface {
	RunMode() RunMode
	SetMode(mode RunMode)
	RunName() string
	LoggerConfig() *LoggerConfig
	RestServerConfig() *RestServerConfig
	GRPCServerConfig() *GRPCServerConfig
	UnmarshalYaml(v interface{}) error
}

type Grouper interface {
	NewGroup(prefix string, middleware ...HandlerFunc) Grouper
	Use(filter ...HandlerFunc)
	Resource(uri string, resource interface{}) Grouper
	OPTIONS(uri string, handler HandlerFunc)
	HEAD(uri string, handler HandlerFunc)
	POST(uri string, handler HandlerFunc)
	GET(uri string, handler HandlerFunc)
	PUT(uri string, handler HandlerFunc)
	PATCH(uri string, handler HandlerFunc)
	DELETE(uri string, handler HandlerFunc)
	Handler(method, uri string, fn HandlerFunc)
	HandlerFunc(method, uri string, fn http.HandlerFunc)
	MountRPC(rpc GRPCService)
}

type RestService interface {
	Init(config Configer, group Grouper)
	Filters()
	Resources()
}

// additional methods for accessing metadata about the service.
type GRPCService interface {
	Register(gs *grpc.Server)
}

type Logger interface {
	RequestID() string
	New(loggerID string) Logger
	Reuse(l Logger)

	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Output interface {
	GetStatus() int
}
