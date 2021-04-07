package middleware

import (
	"easy-go/src/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"runtime"
)

type Context struct {
	*gin.Context
}

func (c *Context) ResponsePanic(msg ...interface{}){
	c.JSON(http.StatusOK,gin.H{
		"status":"0000",
		"msg":"system error",
	})
}


func RouterHandler(routerFunc func(*Context)) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := &Context{
			Context:gctx,
		}
		defer func(){
			if err := recover(); err != nil {
				stack := make([]byte,1024*8)
				stack = stack[:runtime.Stack(stack,false)]
				httpRquest,_ := httputil.DumpRequest(ctx.Request,true)
				logger.GetLogger().Errorf(string(stack))
				logger.GetLogger().Errorf(string(httpRquest))
				ctx.ResponsePanic(err)
			}
		}()

		routerFunc(ctx)
	}
}

