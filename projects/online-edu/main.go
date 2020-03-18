package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"go-learn/projects/online-edu/config"
	"go-learn/projects/online-edu/order/orderRouter"
	"go-learn/projects/online-edu/user/userRouter"

	"net/http"
	"time"
)

func init() {
	flag.StringVar(&config.Mode, "mode", "local", "-mode 'pro'")
	flag.StringVar(&config.ServerName, "server_name", "online_edu_srv", "-server_name 'online_edu_0'")
	flag.StringVar(&config.HttpPort, "http_port", ":8112", "-http_port ':8112'")
	flag.Parse()

	config.InitConfig()
}
func main() {
	go myHTTP()
	select {}
}

func myHTTP() {
	r := gin.New()
	r.Use(gin.Recovery())

	// user
	userRouter.HTTPUserRouter(r)
	// order
	orderRouter.HTTPOrderRouter(r)
	// course
	// courseRouter.HTTPCourseRouter(r)

	var httpPort string
	if config.HttpPort == "" {
		httpPort = config.Cfg.GetString("httpPort")
	}
	s := &http.Server{
		// :8112默认
		Addr:           httpPort,
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}

	fmt.Println("listen on " + httpPort)
	if e := s.ListenAndServe(); e != nil {
		panic(e)
	}
}
