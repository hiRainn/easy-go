package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(ctx *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusOK,gin.H{
				"code":-1,
				"msg": fmt.Sprintf("请求失败：%v",err.Error()),
			})
		} else {
			ctx.Next()
		}
	}()
	traceId := getTraceId()
	ctx.Set("trace_id", traceId)
}

func getTraceId() string {
	var traceId string
	return traceId
}