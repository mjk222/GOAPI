// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"net/http"
	"strings"

	"github.com/GOAPI/routes"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理404请求
	router.NoRoute(func(c *gin.Context) {

		// 获取标头信息的Accept信息
		// Accept描述了客户端希望服务器返回的响应body数据类型
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {

			// 如果是HTML
			c.String(http.StatusNotFound, "页面返回404")
		} else {

			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义,请确认url和请求方法是否正确",
			})
		}

	})
}
