package logger

import (
	"bytes"
	"easy-go/src/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type Logger struct {
	*logrus.Logger
	LogPath string
	TraceId string
	Caller  string
	Expire  int64
	Field   map[string]string
}

var logPoll map[string]*Logger

//
func GetLogger(traceId ...string) *Logger {
	var key string
	var expire int64
	if len(traceId) == 0 {
		key = ""
	} else {
		key = traceId[0]
		expire = time.Now().Unix() + 300
	}
	if logPoll == nil {
		logPoll = make(map[string]*Logger)
	}
	if _, ok := logPoll[key]; !ok {
		logPoll[key] = &Logger{Logger: logrus.New(), TraceId: key, Expire: expire}
		cfg := config.GetConf()
		if cfg.LogConfig.LogFormat == "json" {
			logPoll[key].SetFormatter(&logrus.JSONFormatter{})
		} else {
			logPoll[key].SetFormatter(&logrus.TextFormatter{})
		}
	}

	return logPoll[key]
}

//todo 释放文件句柄
func (l *Logger) Free() {
	/*
	删除句柄代码
	*/
	delete(logPoll, l.TraceId)
}

func Init() {
	var (
	//file *os.File
	//err error
	)
	//path := config.GetConf().LogConfig.LogPath
	//if file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm); err !=nil{
	//	logrus.Error("打开日志文件错误：", err)
	//}
	cfg := config.GetConf()
	logPoll = make(map[string]*Logger)
	logPoll[""] = &Logger{Logger: logrus.New()}

	if cfg.LogConfig.LogFormat == "json" {
		logPoll[""].SetFormatter(&logrus.JSONFormatter{})
	} else {
		logPoll[""].SetFormatter(&logrus.TextFormatter{})
	}

	go manageLogPoll()
}

//release handle of log
func manageLogPoll() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <- ticker.C:
			for _,v := range logPoll {
				if v.Expire < time.Now().Unix() {
					v.Free()
				}
			}
		}
	}
}

//获取请求参数
func getRequestRaw(c *gin.Context) ([]byte, error) {
	var bodyBytes []byte // 我们需要的body内容

	// 从原有Request.Body读取
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, fmt.Errorf("Invalid request body")
	}

	// 新建缓冲区并替换原有Request.body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// 当前函数可以使用body内容
	return bodyBytes, nil
}
