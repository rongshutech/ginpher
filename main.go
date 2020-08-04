/*
@Time: 2020-08-03 10:20
@Auth: xujiang
@File: main.go
@Software: GoLand
@Desc: TODO
*/
package main

import (
	"context"
	"fmt"
	"ginpher/app/config"
	"ginpher/app/middleware/error_handler"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg_env := os.Getenv("cfg_env")
	log.Printf("running with config env as %s.", cfg_env)
	config.InitConfig(cfg_env)

	gin.SetMode(config.ServerConfig.AppMode)
	engine := gin.New()

	//panic 处理
	error_handler.Register(engine)

	// 性能分析 - 正式环境不要使用！！！
	if config.ServerConfig.AppMode != "release" {
		pprof.Register(engine)
	}

	log.Printf("%v", config.ServerConfig)

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", config.ServerConfig.AppHost,
			config.ServerConfig.AppPort),
		Handler:      engine,
		ReadTimeout:  time.Duration(config.ServerConfig.AppReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ServerConfig.AppWriteTimeout) * time.Second,
	}

	log.Println("|  github.com/rongshutech/ginpher   |")
	log.Println("|-----------------------------------|")
	log.Println("|  Go Http Server Start Successful  |")
	log.Println("|    Port" + fmt.Sprintf(" %d", config.ServerConfig.AppPort) +
		"     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "       |")
	log.Println("|-----------------------------------|")
	log.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
