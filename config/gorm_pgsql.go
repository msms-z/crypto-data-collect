package config

//type GeneralDB struct {
//	Path         string `json:"path" yaml:"path"`                     // 服务器地址:端口
//	Port         string `json:"port" yaml:"port"`                     //:端口
//	Config       string `json:"config" yaml:"config"`                 // 高级配置
//	Dbname       string `json:"db-name" yaml:"db-name"`               // 数据库名
//	Username     string `json:"username" yaml:"username"`             // 数据库用户名
//	Password     string `json:"password" yaml:"password"`             // 数据库密码
//	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
//	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
//	LogMode      string `json:"log-mode" yaml:"log-mode"`             // 是否开启Gorm全局日志
//	LogZap       bool   `json:"log-zap" yaml:"log-zap"`               // 是否通过zap写入日志文件
//}
//
//type Pgsql struct {
//	GeneralDB `yaml:",inline" mapstructure:",squash"`
//}
