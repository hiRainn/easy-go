package logger

import (
	"bytes"
	"easy-go/src/config"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type Logger struct {
	*logrus.Logger
	LogPath string
	TraceId string
	Caller  map[string]interface{}
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
		logPoll[key].Logger.SetLevel(logrus.DebugLevel)
		cfg := config.GetConf()
		if cfg.LogConfig.LogFormat == "json" {
			logPoll[key].SetFormatter(&logrus.JSONFormatter{})
		} else {
			logPoll[key].SetFormatter(&logrus.TextFormatter{})
		}
		setOutput(logPoll[key])
	}

	return logPoll[key]
}

//free
func (l *Logger) Free() {
	logPoll[l.TraceId].Logger = nil
	delete(logPoll, l.TraceId)
}

func Init() {

	cfg := config.GetConf()
	logPoll = make(map[string]*Logger)
	//new log with key "" for log pool
	logPoll[""] = &Logger{Logger: logrus.New()}
	logPoll[""].Logger.SetLevel(logrus.DebugLevel)
	if cfg.LogConfig.LogFormat == "json" {
		logPoll[""].SetFormatter(&logrus.JSONFormatter{})
	} else {
		logPoll[""].SetFormatter(&logrus.TextFormatter{})
	}
	setOutput(logPoll[""])
	go manageLogPoll()
}

//release handle of log
func manageLogPoll() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <- ticker.C:
			for _,v := range logPoll {
				//filter trace_id eq ""
				if v.Expire < time.Now().Unix() && v.TraceId != "" {
					v.Free()
				}
			}
		}
	}
}

func setOutput(log *Logger) {
	cfg := config.GetConf()
	logPath := cfg.LogConfig.LogPath
	if logPath[len(logPath)-1:] != "/" {
		logPath = logPath + "/"
	}
	writer, err := rotatelogs.New(
		logPath + "%Y%m%d.log",
		rotatelogs.WithMaxAge(time.Duration(cfg.LogConfig.SaveDay)*time.Hour * 24),
		)
	if err != nil {
		log.Fatalf("启动日志失败：%v", err.Error())
	}
	log.Logger.SetOutput(writer)
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
