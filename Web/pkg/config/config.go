package config

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	SERVER_PORT = "server.port"
	SERVER_IP = "server.ip"

	PROFILE_ACTIVE = "profile.active"

	TEMPLATE_PATH   = "template.path"
	TEMPLATE_SUFFIX = "template.suffix"

	FILE_PATH = "file.path"

	STATIC_PATH = "static.path"
)



// 读取配置
func InitConfig() {
	// 初始环境
	viper.SetConfigName("application")
	viper.AddConfigPath("conf/")
	viper.SetConfigType("yaml")
	if e := viper.ReadInConfig(); e != nil {
		fmt.Println(e.Error())
	}

	// 根据环境追加配置
	active := viper.GetString(PROFILE_ACTIVE)
	if active != "" {
		viper.SetConfigName(fmt.Sprintf("application-%s", active))
		viper.MergeInConfig()
	}

	_, cancel := context.WithCancel(context.Background());
	go func() {
		viper.OnConfigChange(func(in fsnotify.Event) {
			InitConfig()
			cancel()
			fmt.Println(".yaml changed")
		})
		viper.WatchConfig()
	}()
}