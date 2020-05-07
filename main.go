package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/x554462/demo/app"
	"github.com/x554462/demo/app/router"
	"github.com/x554462/demo/middleware/mango/library/conf"
	"log"
	"net/http"
	"syscall"
)

func main() {

	app.Setup()

	gin.SetMode(conf.ServerConf.RunMode)

	endless.DefaultReadTimeOut = conf.ServerConf.ReadTimeout
	endless.DefaultWriteTimeOut = conf.ServerConf.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	var handler http.Handler
	handler = route()
	if conf.ServerConf.HttpTimeout != 0 {
		handler = http.TimeoutHandler(handler, conf.ServerConf.HttpTimeout, "timeout")
	}
	server := endless.NewServer(conf.ServerConf.Addr, handler)
	server.BeforeBegin = func(addr string) {
		log.Printf("start http server listening %s, pid is %d", addr, syscall.Getpid())
	}

	server.ListenAndServe()
}

func route() *gin.Engine {
	r := gin.New()
	if gin.IsDebugging() {
		r.Use(gin.Logger())
		pprof.Register(r)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.InitRouter(r)
	return r
}
