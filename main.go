package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/bootstrap"
)

func main() {
	router := gin.New()
	bootstrap.SetupRoute(router)
	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
