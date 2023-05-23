package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/bootstrap"
	btsConfig "go-blog/config"
	"go-blog/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	router := gin.New()
	bootstrap.SetupDB()
	bootstrap.SetupRoute(router)
	bootstrap.SetupLogger()
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
