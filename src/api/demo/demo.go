package demo

import (
	"easy-go/src/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

type Demo struct{}

func(t *Demo) Router(router *gin.RouterGroup) {
	router.GET("hello", middleware.RouterHandler(t.Hello))
}

func (t *Demo) Hello(ctx *middleware.Context) {
	ctx.ResponseSuccess(map[string]interface{}{"id":10,"name":"heihei"},map[string]interface{}{"trace_id":"testid18d7adsa","time":time.Now()})
}