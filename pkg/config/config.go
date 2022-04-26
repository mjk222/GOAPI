package config

import (
	"fmt"
	"os"

	"github.com/GOAPI/pkg/helpers"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

// viper 库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 在动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {

	// 1.初始化 Viper 库
	viper = viperlib.New()
	// 2.配置类型，支持"json","toml","yaml","yml","properties","props","prop","env","dotenv"
	viper.SetConfigType("env")
	// 3.环境变量配置文件查找的路径，相对于main.go
	viper.AddConfigPath(".")
	// 4.设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5.读取环境变量（支持flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1.加载环境变量
	loadEnv(env)
	// 2.注册配置信息
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// 加载 .env 文件
func loadEnv(envSuffix string) {

	// 默认加载 .env 文件，如果有传参 -env name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filepath
		}
	}

	fmt.Println("envPath is ", envPath)
	// 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err is %v\n", err)
		panic(err)
	}

	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetBool获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}
