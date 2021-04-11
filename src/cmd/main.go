package main

import (
	"easy-go/src/api"
	"easy-go/src/config"
	"easy-go/src/logger"
	"flag"
	"fmt"
	"log"
)

func main() {
	initConf()
	go api.StartApiServer(config.Conf.ServerPort)
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
	if err = config.InitConf(*cfg); err != nil {
		log.Fatalf("init config error：%v",err.Error())
	}
	logger.Init()
	logger.GetLogger().Info("logger start...")
	fmt.Println("gelodsa")
}

