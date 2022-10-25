package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-base/common"
	"go-api-base/common/initialize"
	"go-api-base/config"
	"go-api-base/middleware"
	"go-api-base/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	initialize.Config()
	initialize.Mysql()
	initialize.Log()
	r := gin.New()
	r.Use(middleware.Logger(common.Logger, time.RFC3339, true), middleware.Recovery(common.Logger, true))
	r.Use(middleware.Cors())
	router.Router(r)
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		go func() {
			common.Logger.Sync()
		}()
		if err := server.Close(); err != nil {
			log.Fatal("服务关闭出现错误:", err)
		}
	}()
	log.Println("正在启动服务器:", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("服务启动错误:", err)
	}

	log.Println("服务器已退出")
}
