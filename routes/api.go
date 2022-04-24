package routes

import (
	"net/http"

	"github.com/GOAPI/app/http/controllers/api/v1/auth"
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

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已经注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否已经注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}
}
