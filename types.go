package mpc

import (
	"net/http"
	"sync"
)

type Context struct {
	Response *Response
	Request  *http.Request
	mux      sync.RWMutex
	index    int8
	handlers []HandlerFunc
}

type Response struct {
	http.ResponseWriter

	mux    sync.RWMutex
	status int
	code   int
}

type HandlerFunc func(ctx *Context)

type Configer interface {
	RunMode() RunMode
	RunName() string
	SetMode(mode RunMode)
	UnmarshalYaml(v interface{}) error
}

type Grouper interface {
	NewGroup(prefix string, middleware ...HandlerFunc) Grouper
	Use(filter ...HandlerFunc)
	Resource(uri string, resource interface{}) Grouper
	OPTIONS(uri string, handler HandlerFunc)
	//HEAD(uri string, handler Middleware)
	//POST(uri string, handler Middleware)
	GET(uri string, handler HandlerFunc)
	//PUT(uri string, handler Middleware)
	//PATCH(uri string, handler Middleware)
	//DELETE(uri string, handler Middleware)
	//Any(uri string, handler Middleware)
	Handler(method, uri string, fn HandlerFunc)
	HandlerFunc(method, uri string, fn http.HandlerFunc)
	//Handler(method, uri string, handler http.Handler)
	//Handle(method, uri string, handler Middleware)
	////MountRPC(method string, rpc RPCService)
	//MockHandle(method, uri string, recorder http.ResponseWriter, handler Middleware)
}

// A Service represents RESTful service interface of pedestal.
type RestService interface {
	Init(config Configer, group Grouper)
	Filters()
	Resources()
}

// additional methods for accessing metadata about the service.
type RPCService interface {
	Version() string
	ServiceNames() []string
	ServiceRegistry(prefix string) []ServiceHandler
	ServiceDescriptor() ([]byte, int)
}

type ServiceHandler struct {
	Method  string
	URI     string
	Handler HandlerFunc
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
