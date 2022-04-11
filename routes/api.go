package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			// JSON格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "majikun",
			})
		})
	}
}
