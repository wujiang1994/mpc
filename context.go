package mpc

import "net/http"

func (c *Context) Next() {
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
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

func (c *Context) AddHeader(key, value string)  {
	c.Response.Header().Add(key, value)
}

func (c *Context) SetStatus(status int) {
	c.Response.status = status
}

func (c *Context) Json(data interface{}) error {
	return c.Return(NewJsonRender(c.Response), data)
}

func (c *Context) Return(rr Render, data interface{}) error {
	if data == nil {
		return nil
	}
	c.SetHeader("Content-Type", rr.ContentType())
	c.SetStatus(http.StatusCreated)
	return rr.Render(data)
}