package routes

import (
	"GFile/pkg/core"
	fileService "GFile/pkg/service"
	"net/http"
	"strconv"
)

/**
 获取文件内容
 */
func GetFile(w http.ResponseWriter, r *http.Request)  {
	data := fileService.QueryAll()
	write(w, r, core.Success(data))
}

/**
 分页获取
 */
func Page(w http.ResponseWriter, r *http.Request)  {
	size := 10
	page := 1
	var classify string
	if v, ok := r.URL.Query()["size"]; ok && len(v[0]) > 0 {
		size, _ = strconv.Atoi(v[0])
	}
	if v, ok := r.URL.Query()["page"]; ok && len(v[0]) > 0 {
		page, _ = strconv.Atoi(v[0])
	}
	if v, ok := r.URL.Query()["classify"]; ok && len(v[0]) > 0 && v[0] != "所有" {
		classify = v[0]
	}

	data := fileService.Page(page, size, classify)
	write(w, r,  core.Success(data))
}