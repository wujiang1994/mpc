package mpc

import "fmt"

func CORS() HandlerFunc {
	return func(ctx *Context) {
		ctx.Response.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header().Set("Access-Control-Max-Age", "600")
		ctx.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		ctx.Response.Header().Set("Access-Control-Allow-Headers", "token,Keep-Alive,User-Agent,X-Requested-With,Cache-Control,Content-Type,Authorization")
		ctx.Next()
	}
}

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("internal panic: %v", err)
				_, _ = ctx.Response.Write([]byte(errMsg))
			}
		}()
	}
}
