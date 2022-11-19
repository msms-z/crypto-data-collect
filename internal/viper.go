package internal

import (
	"crypto-data-collect/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {

	var currentFilePath string

	if len(path) == 0 {
		//flag.StringVar(&config, "c", "", "choose config file.")
		//flag.Parse()
		currentFilePath = ConfigEnv
	} else {
		currentFilePath = path[0]
	}
	fmt.Printf("正在使用的Config路径为%s\n", currentFilePath)

	v := viper.New()
	v.SetConfigFile(currentFilePath)
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n ", err))
	}

	if err := v.Unmarshal(&global.SERVER_CONFIG); err != nil {
		panic(fmt.Errorf("Error read config file:  %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)

	})

	return v
}
