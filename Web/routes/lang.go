package routes

import (
	"GFile/pkg/config"
	"GFile/pkg/core"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"html/template"
	"net/http"
)

func write(w http.ResponseWriter, r *http.Request, result *core.Result) {
	origin := r.Header.Get("Origin")
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", origin)
	header.Add("Access-Control-Allow-Credentials", "true");
	header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT");
	header.Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With");
	header.Add("Access-Control-Expose-Headers", "*");
	w.Write(result.ToByte())
}

func render(w http.ResponseWriter, tpl string, data interface{}) {
	path := viper.GetString(config.TEMPLATE_PATH)
	suffix := viper.GetString(config.TEMPLATE_SUFFIX)

	if tpl, e := template.ParseFiles(fmt.Sprint(path, tpl, ".", suffix)); e == nil {
		tpl.Execute(w, data)
		return
	}
	panic(errors.New("模版解析错误"))
}