package mpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type Context struct {
	Response *Response
	Request  *Request
	mux      sync.RWMutex
	handlers []HandlerFunc
}

type Request struct {
	*http.Request

	requestID string
	Prams     *Prams
	Logger    Logger
}

type Prams struct {
	*http.Request
	rawBody []byte
	rawErr  error
}

type Response struct {
	http.ResponseWriter

	mux    sync.RWMutex
	status int
	code   int
}

func (c *Context) Do() {
	for _, handler := range c.handlers {
		handler(c)
	}
}

func (c *Context) HasHeader(key string) bool {
	_, ok := c.Request.Header[http.CanonicalHeaderKey(key)]
	return ok
}

func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Context) SetHeader(key, value string) {
	c.Response.Header().Set(key, value)
}

func (c *Context) AddHeader(key, value string) {
	c.Response.Header().Add(key, value)
}

func (c *Context) SetStatus(status int) {
	c.Response.status = status
}

func (c *Context) Json(data Output) error {
	return c.Return(NewJsonRender(c.Response), data)
}

func (c *Context) Return(rr Render, data Output) error {
	if data == nil {
		return nil
	}
	c.SetHeader("Content-Type", rr.ContentType())
	c.SetStatus(http.StatusOK)
	if data.GetStatus() > 0 {
		c.SetStatus(data.GetStatus())
	}
	return rr.Render(data)
}

func NewPrams(r *http.Request) *Prams {
	return &Prams{
		Request: r,
	}
}

func (p *Prams) Get(key string) string {
	return p.URL.Query().Get(key)
}

func (p *Prams) Json(v interface{}) error {
	p.rawBody, p.rawErr = ioutil.ReadAll(p.Body)

	if p.rawErr != nil {
		return p.rawErr
	}
	p.Body.Close()
	p.Body = ioutil.NopCloser(bytes.NewReader(p.rawBody))

	return json.Unmarshal(p.rawBody, v)
}
