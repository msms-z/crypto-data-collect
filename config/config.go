package config

import "fmt"

type Config struct {
	//Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type Redis struct {
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 数据库名
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 数据库端口号
	Dbname   int    `mapstructure:"dbname" json:"dbname" yaml:"dbname"`       // 数据库端口号
}

type Database struct {
	Adopter      string `mapstructure:"adopter" json:"adopter" yaml:"adopter"`                      // 服务器地址:端口
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 数据库名
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口号
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         //全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
}

func (db *Database) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		db.Host, db.Username, db.Password, db.Dbname, db.Port)
}
