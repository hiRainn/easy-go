package middleware

import (
	"easy-go/src/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"reflect"
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

func (c *Context) ResponseSuccess(data ...interface{}) {
	response := gin.H{
		"code":0,
		"msg":"success",
	}
	if len(data) > 0 {
		response["data"] = data[0]
	}
	if len(data) > 1 {
		if reflect.TypeOf(data[1]).Kind() == reflect.Map && reflect.TypeOf(data[1]).Key().Kind() == reflect.String {
			for _,key := range reflect.ValueOf(data[1]).MapKeys() {
				if key.String() != "code" && key.String() != "msg" && key.String() != "data" {
					response[key.String()] = reflect.ValueOf(data[1]).MapIndex(key).Interface()
				}
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

