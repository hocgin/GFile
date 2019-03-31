package main

import (
	"Web/src/config"
	"Web/src/routes"
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"net/http"
)

func main() {

	initConfig()

	initHandle()

	port := viper.GetString(constant.SERVER_PORT)

	// 静态文件服务器

	static := viper.GetString(constant.STATIC_PATH)
	http.Handle("/", http.FileServer(http.Dir(static)))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// 读取配置
func initConfig() {
	// 初始环境
	viper.SetConfigName("application")
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	if e := viper.ReadInConfig(); e != nil {
		fmt.Println(e.Error())
	}

	// 根据环境追加配置
	active := viper.GetString(constant.PROFILE_ACTIVE)
	if active != "" {
		viper.SetConfigName(fmt.Sprintf("application-%s", active))
		viper.MergeInConfig()
	}

	_, cancel := context.WithCancel(context.Background());
	go func() {
		viper.OnConfigChange(func(in fsnotify.Event) {
			initConfig()
			cancel()
			fmt.Println(".yaml changed")
		})
		viper.WatchConfig()
	}()
}

// 初始化控制器
func initHandle() {
	//http.Handle()
	http.HandleFunc("/index", routes.Index)
	http.HandleFunc("/file", routes.GetFile)
	http.HandleFunc("/file/_search", routes.Page)
	http.HandleFunc("/analysis/video", routes.AnalysisVideo)
	http.HandleFunc("/analysis/file", routes.AnalysisFile)
	http.HandleFunc("/classify", routes.QueryAllPath)
}
