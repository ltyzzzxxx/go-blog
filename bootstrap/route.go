package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-blog/app/http/middlewares"
	"go-blog/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)
	routes.RegisterAPIRoutes(router)
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		//gin.Logger(),
		middlewares.Logger(),
		//gin.Recovery(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
