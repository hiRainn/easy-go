package logger

import (
	"bytes"
	"easy-go/src/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var Log = logrus.New()

func Init(cfg *config.Config) {
	if cfg.LogConfig.LogFormat == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{})
	}
	
}

//获取请求参数
func getRequestRaw(c *gin.Context) ([]byte,error) {
	var bodyBytes []byte // 我们需要的body内容

	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, fmt.Errorf("Invalid request body")
	}

	// 新建缓冲区并替换原有Request.body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// 当前函数可以使用body内容
	return bodyBytes,nil
}
