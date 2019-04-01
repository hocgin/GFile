package routes

import (
	"GFile/pkg/config"
	"GFile/pkg/core"
	"github.com/spf13/viper"
	"net/http"
)

func QueryAllPath(w http.ResponseWriter, r *http.Request) {
	paths := viper.GetStringMapString(config.FILE_PATH)
	write(w, r, core.Success(paths))
}
