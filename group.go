package mpc

import (
	"github.com/rs/xid"
	"mpc/utils"
	"net/http"
	"path"
	"sync"
)

type AppGroup struct {
	mux        sync.Mutex
	server     *AppServer
	prefix     string
	middleware []HandlerFunc
	route      map[string]map[string]HandlerFunc
}

func NewAppGroup(prefix string, appServer *AppServer) *AppGroup {
	return &AppGroup{
		server:     appServer,
		prefix:     prefix,
		middleware: []HandlerFunc{},
		route:      make(map[string]map[string]HandlerFunc),
	}
}

func (a *AppGroup) NewGroup(prefix string, middleware ...HandlerFunc) Grouper {
	a.mux.Lock()
	defer a.mux.Unlock()

	return &AppGroup{
		server:     a.server,
		prefix:     a.buildPrefix(prefix),
		middleware: a.buildHandlers(middleware...),
		route:      make(map[string]map[string]HandlerFunc),
	}
}

func (a *AppGroup) Use(middleware ...HandlerFunc) {
	a.middleware = append(a.middleware, middleware...)
}

func (a *AppGroup) Resource(uri string, resource interface{}) Grouper {
	return nil
}

func (a *AppGroup) OPTIONS(uri string, handler HandlerFunc) {
	a.Handler(http.MethodOptions, uri, handler)
}

func (a *AppGroup) HEAD(uri string, handler HandlerFunc) {
	a.Handler(http.MethodHead, uri, handler)
}

func (a *AppGroup) POST(uri string, handler HandlerFunc) {
	a.Handler(http.MethodPost, uri, handler)
}

func (a *AppGroup) GET(uri string, handler HandlerFunc) {
	a.Handler(http.MethodGet, uri, handler)
}

func (a *AppGroup) PUT(uri string, handler HandlerFunc) {
	a.Handler(http.MethodPut, uri, handler)
}

func (a *AppGroup) PATCH(uri string, handler HandlerFunc) {
	a.Handler(http.MethodPatch, uri, handler)
}

func (a *AppGroup) DELETE(uri string, handler HandlerFunc) {
	a.Handler(http.MethodDelete, uri, handler)
}

func (a *AppGroup) Handler(method, uri string, handler HandlerFunc) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if len(a.server.route[a.buildPrefix(uri)]) == 0 {
		a.server.route[a.buildPrefix(uri)] = make(map[string]HandlerFunc)
	}
	a.server.route[a.buildPrefix(uri)][method] = handler
}

func (a *AppGroup) HandlerFunc(method, uri string, handler http.HandlerFunc) {
}

func (a *AppGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := new(Context)
	id := xid.New().String()
	c.Request = &Request{
		Request:   r,
		requestID: id,
		Prams:     NewPrams(r),
		Logger:    globalLogger.New(id),
	}
	c.Response = &Response{
		ResponseWriter: w,
	}
	if len(a.route[c.Request.URL.Path]) == 0 {
		NotFoundHandler.ServeHTTP(w, r)
		return
	}
	if a.route[c.Request.URL.Path][c.Request.Method] == nil {
		MethodNotAllowedHandler.ServeHTTP(w, r)
		return
	}
	c.handlers = append(c.handlers, a.middleware...)
	c.handlers = append(c.handlers, a.route[c.Request.URL.Path][c.Request.Method])
	c.Do()
	return
}

func (a *AppGroup) buildPrefix(suffix string) (prefix string) {
	utils.Assert(suffix[0] == '/', "path must begin with '/'")
	if len(suffix) == 0 {
		prefix = a.prefix
		return
	}
	prefix = path.Join(a.prefix, suffix)
	return
}

func (a *AppGroup) buildHandlers(handlers ...HandlerFunc) []HandlerFunc {
	a.middleware = append(a.middleware, handlers...)
	return a.middleware
}
