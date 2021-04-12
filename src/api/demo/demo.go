package demo

import (
	"easy-go/src/middleware"
	"github.com/gin-gonic/gin"
)

type Demo struct{}

func(t *Demo) Router(router *gin.RouterGroup) {
	router.GET("hello", middleware.RouterHandler(t.Hello))
}

func (t *Demo) Hello(ctx *middleware.Context) {
	ctx.ResponseSuccess(nil,map[string]interface{}{"trace_id":"dsak1932kdasd",})
}