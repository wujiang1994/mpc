package mpc

import (
	"encoding/json"
	"net/http"
)

type Render interface {
	ContentType() string
	Render(v interface{}) error
}

type JsonRender struct {
	w http.ResponseWriter
}

func NewJsonRender(w http.ResponseWriter) Render {
	return &JsonRender{w}
}

func (j *JsonRender) ContentType() string {
	return "application/json"
}

func (j *JsonRender) Render(v interface{}) error {
	if v == nil {
		return nil
	}
	return json.NewEncoder(j.w).Encode(v)
}
