package main

import (
	"easy-go/etc"
	"easy-go/src/api"
	"easy-go/src/logger"
	"flag"
	"log"
)

func main() {
	initConf()
	go api.StartApiServer(etc.Conf.ServerPort)
	select {}
}


func initConf() {
	var err error
	//fileName := "../../etc/config.yaml"
	//dev path
	fileName := "D:\\Go\\easy-go\\etc\\config.yaml"
	//fileName := "/Applications/MAMP/gopath/easy-go/etc/config.yaml"
	cfg := flag.String("c",fileName,"Config file")
	flag.Parse()
	if err = etc.InitConf(*cfg); err != nil {
		log.Fatalf("init config error：%v",err.Error())
	}
	logger.Init()
	logger.GetLogger().Info("logger start...")
}

