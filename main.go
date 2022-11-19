package main

import (
	"crypto-data-collect/global"
	"crypto-data-collect/internal"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	fmt.Println("开始监听获取数据")

	internal.Viper()

	global.SERVER_REDIS = internal.Redis()

	global.SERVER_LOG = internal.InitZap()

	global.SERVER_DB = internal.GormPgSql(global.SERVER_CONFIG.Database)
	if global.SERVER_DB != nil {

		internal.RegisterTables(global.SERVER_DB)

		db, _ := global.SERVER_DB.DB()

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.SERVER_LOG.Error("Close Database Error: {}", zap.Error(err))
			}
		}(db)
	}

}
