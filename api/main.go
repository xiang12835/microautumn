package main

import (
	"context"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	//"runtime"
	_ "github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	_ "gopkg.in/redis.v4"
	_ "microautumn/api/docs"
	"microautumn/api/routers"
	_ "microautumn/models"
)

func main() {

	fmt.Println("[Server Starting]...")

	gin.SetMode(gin.ReleaseMode)
	router := routers.InitRouter()

	server := endless.NewServer("127.0.0.1:8080", router)
	server.ReadTimeout = 3 * time.Second
	server.WriteTimeout = 3 * time.Second

	server.BeforeBegin = func(add string) {
		log.Printf(add)
	}

	go func() {
		// 监听请求
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅Shutdown（或重启）服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
	}
	log.Println("Server exiting")

}
