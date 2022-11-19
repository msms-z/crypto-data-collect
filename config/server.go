package config

type ServerConfig struct {
	//Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
}
