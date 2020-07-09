package controllers

import (
	"mpc"
	"mpc/testapp/app/errors"
	"mpc/testapp/app/services"
	"net/http"
)

type Application struct {
	v1 mpc.Grouper
	v2 mpc.Grouper
}

func New() mpc.RestService {
	return &Application{}
}

func (a *Application) Init(config mpc.Configer, group mpc.Grouper) {
	services.SetupMpcTest()
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
		id := ctx.Request.Prams.Get("id")
		log := mpc.NewAppLogger()
		log.Infof("test for logger, id: %s", id)
		ctx.Json(NewCommonRespFromErr(errors.ErrForTest))
	})
	a.v2.Handler(http.MethodPost, "/test", func(ctx *mpc.Context) {
		req := new(services.MpcTestDoReq)
		if err := ctx.Request.Prams.Json(req); err != nil {
			ctx.Json(NewCommonRespFromErr(err))
			return
		}
		data, err := services.MpcTest.Do(ctx, req)
		if err != nil {
			ctx.Json(NewCommonRespFromErr(err))
			return
		}
		ctx.Json(NewCommonRespFromData(data))
	})

	v3 := a.v2.NewGroup("/v3")
	v3.Use(mpc.Recovery())
	v3.GET("/test", func(ctx *mpc.Context) {
		data := struct {
			A int
		}{
			A: 555,
		}
		ctx.Json(NewCommonRespFromData(data))
	})
}

type CommonResp struct {
	Code    int         `json:"code"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *CommonResp) GetStatus() int {
	return c.Status
}

func NewCommonRespFromData(data interface{}) *CommonResp {
	return &CommonResp{
		Data: data,
	}
}

func NewCommonRespFromErr(err error) *CommonResp {
	return &CommonResp{
		Code:    -1,
		Message: err.Error(),
	}
}
