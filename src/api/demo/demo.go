package demo

import (
	"easy-go/src/logger"
	"easy-go/src/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

type Demo struct{}

func(t *Demo) Router(router *gin.RouterGroup) {
	router.GET("hello", middleware.RouterHandler(t.Hello))
}

func (t *Demo) Hello(ctx *middleware.Context) {
	logger.GetLogger("asdasdsadsa").Error("logger start12311...")
	logger.GetLogger("asdasdsadsa").Debug("logger start11111...")
	ctx.ResponseSuccess(map[string]interface{}{"id":10,"name":"heihei"},map[string]interface{}{"trace_id":"testid18d7adsa","time":time.Now()})
}