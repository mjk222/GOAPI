package main

import (
	"flag"
	"fmt"

	"github.com/GOAPI/bootstrap"
	btsConfig "github.com/GOAPI/config"
	"github.com/GOAPI/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，以来命令行  -env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
	fmt.Println("finish")
}
