package main

import (
	"GFile/pkg/config"
	"GFile/routes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	config.InitConfig()

	initHandle()

	port := viper.GetString(config.SERVER_PORT)

	// 静态文件服务器
	static := viper.GetString(config.STATIC_PATH)
	http.Handle("/", http.FileServer(http.Dir(static)))

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
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
