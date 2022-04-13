package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"github.com/GOAPI/pkg/config"
	"github.com/GOAPI/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=local", config.Get("database.mysql.username"), config.Get("database.mysql.password"), config.Get("databasse.mysql.host"), config.Get("database.mysql.post"), config.Get("database.mysql.database"), config.Get("database.mysql.charset"))
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqliter":
		// 初始化 sqlite
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))

	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.nysql.max_idel_connections"))

	// 设置每个连接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")))
}
