package local

import (
	"Web/src/config"
	"Web/src/core/result"
	"Web/src/model"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

// 查询所有
func QueryAll() []*model.FileInfo {
	return all()
}

// 分页查询
func Page(page int, size int, classify string) []*model.FileInfo {
	var data []*model.FileInfo
	if classify == "" {
		data = all()
	} else {
		data = dir(classify)
	}
	start, end := result.WrapPage(page, size, len(data))
	return data[start:end]
}

// 解码
func Decode(ePath string) (string, bool) {
	if path, e := url.QueryUnescape(ePath); e == nil {
		return path, true
	}
	return "", false
}

// 编码
func Encode(path string) (string) {
	return url.QueryEscape(path)
}

// 解析标识为文件
func AnalysisLocalFile(ePath string) (*os.File, bool) {
	if path, ok := Decode(ePath); ok {
		if file, err := os.Open(path); err == nil {
			return file, true
		}
	}
	return nil, false
}

// 获取所有文件
func all() ([]*model.FileInfo) {
	paths := viper.GetStringMapString(constant.FILE_PATH)
	var data []*model.FileInfo
	for tag, path := range paths {
		data = append(data, list(tag, path)...)
	}
	return data
}

// 根据 tag 获取文件
func dir(tag string) ([]*model.FileInfo) {
	path := viper.GetStringMapString(constant.FILE_PATH)[tag]
	return list(tag, path)
}

// 获取文件列表
func list(tag string, path string) ([]*model.FileInfo) {
	var rs []*model.FileInfo
	if files, e := ioutil.ReadDir(path); e == nil {
		for _, file := range files {
			if file.IsDir() {
				rs = append(rs, list(tag, filepath.Join(path, file.Name()))...)
			} else {
				info := model.NewFileInfo(file, filepath.Join(path, file.Name()))
				info.Tags = append(info.Tags, tag)
				rs = append(rs, info)
			}
		}
	}
	return rs
}
