package routes

import (
	"Web/src/config"
	"Web/src/core/result"
	"github.com/spf13/viper"
	"net/http"
)

func QueryAllPath(w http.ResponseWriter, r *http.Request) {
	paths := viper.GetStringMapString(constant.FILE_PATH)
	write(w, r, result.Success(paths))
}
