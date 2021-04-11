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
		"code":-1,
		"msg":"system error",
	})
}

func (c *Context) ResponseFail(msg interface{}) {
	c.JSON(http.StatusOK,gin.H{
		"code":1,
		"msg":msg,
	})
}

func (c *Context) ResponseError(msg interface{}) {
	c.JSON(http.StatusOK,gin.H{
		"code":99,
		"msg":msg,
	})
}

func (c *Context) ResponseSuccess(data interface{},key ...map[string]interface{}) {
	response := gin.H{
		"code":0,
		"msg":"success",
		"data":data[0],
	}
	if len(key) > 0 {
		for k,v := range key[0] {
			if k != "code" && k != "msg" && k != "data" {
				response[k] = v
			}
		}
	}
	c.JSON(http.StatusOK,response)
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

