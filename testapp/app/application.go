package app

import (
	"net/http"
	"mpc"
)

type Application struct {
	v1 mpc.Grouper
	v2 mpc.Grouper
}

func New() mpc.RestService {
	return &Application{}
}

func (a *Application) Init(config mpc.Configer, group mpc.Grouper) {
	a.v1 = group.NewGroup("/v1")
	a.v2 = group.NewGroup("/v2")
}

func (a *Application) Filters() {
	a.v1.Use(mpc.CORS())
	a.v2.Use(mpc.Recovery())
}

func (a *Application) Resources() {
	a.Routes()
}

func (a *Application) Routes() {
	a.v1.Handler(http.MethodGet, "/test", func(ctx *mpc.Context) {
		id := ctx.Request.URL.Query().Get("id")
		log := mpc.NewAppLogger()
		log.Infof("你好你好你好你好, %+v", id)
		ctx.Response.Write([]byte("hello world"))
	})
	a.v2.Handler(http.MethodPost, "/test", func(ctx *mpc.Context) {
		ctx.Response.Write([]byte("hello group v2"))
	})

	v3 := a.v2.NewGroup("/v3")
	v3.Use(mpc.Recovery())
	v3.GET("/test", func(ctx *mpc.Context) {
		ctx.Response.Write([]byte("hello group v3"))
	})
}
