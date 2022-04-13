// Package config 站点配置信息
package config

import (
	"fmt"

	"github.com/GOAPI/pkg/config"
)

func init() {
	fmt.Println("config info")
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("APP_NAME", "Gohub"),

			// 当前环境，用以区分多环境，一般为local，stage，production，test
			"env": config.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env("APP_PORT", "3000"),

			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// 设置时区，JWT 会使用，日记记录也会用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
