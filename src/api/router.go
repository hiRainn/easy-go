package api

import (
	"context"
	t "easy-go/src/api/test"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func registerRouter(router *gin.Engine) {
	api := router.Group("/api")
	//if you need different versions
	v1 := api.Group("/v1")

	//test module
	test := v1.Group("/test")
	new(t.Test).Router(test)
}

func StartApiServer(port int) {
	router := gin.Default()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start api server fail：%v", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shuting down server")
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shut down :%v", err.Error())
	}
	log.Println("server exit")

}
