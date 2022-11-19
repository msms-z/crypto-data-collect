package internal

import (
	"crypto-data-collect/config"
	"crypto-data-collect/global"
	"crypto-data-collect/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.ContractData{},
	)
	if err != nil {
		global.SERVER_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.SERVER_LOG.Info("Register table success.")
}

func GormPgSql(dbConfig config.Database) *gorm.DB {
	if dbConfig.Dbname == "" {
		panic("Please check config file, database name cannot be null.")
	}

	// postgresql 配置
	pgConfig := postgres.Config{
		DSN:                  dbConfig.Dsn(),
		PreferSimpleProtocol: false,
	}
	// gorm 配置信息
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConfig.Prefix,
			SingularTable: dbConfig.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if db, err := gorm.Open(postgres.New(pgConfig), &gormConfig); err != nil {
		panic(err)
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConns)
		return db
	}

}
