package middleware

import "github.com/gin-gonic/gin"

func Auth(ctx *gin.Context) {

	traceId := getTraceId()
	ctx.Set("trace_id", traceId)
}

func getTraceId() string {
	var traceId string

	return traceId
}