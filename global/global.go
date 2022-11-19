package global

import (
	"crypto-data-collect/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	// SERVER_CONFIG 配置全局config数据
	SERVER_CONFIG *config.Config
	// SERVER_REDIS 配置全局 Redis Client
	SERVER_REDIS *redis.Client
	SERVER_DB    *gorm.DB
	SERVER_LOG   *zap.Logger

	lock sync.RWMutex
)
