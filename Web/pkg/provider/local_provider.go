package provider

import (
	"GFile/pkg/config"
	"GFile/pkg/util"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

// 查询所有
func QueryAll() []*FileInfo {
	return all()
}

// 分页查询
func Page(page int, size int, classify string) []*FileInfo {
	var data []*FileInfo
	if classify == "" {
		data = all()
	} else {
		data = dir(classify)
	}
	start, end := util.WrapPage(page, size, len(data))
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
func all() ([]*FileInfo) {
	paths := viper.GetStringMapString(config.FILE_PATH)
	var data []*FileInfo
	for tag, path := range paths {
		data = append(data, list(tag, path)...)
	}
	return data
}

// 根据 tag 获取文件
func dir(tag string) ([]*FileInfo) {
	path := viper.GetStringMapString(config.FILE_PATH)[tag]
	return list(tag, path)
}

// 获取文件列表
func list(tag string, path string) ([]*FileInfo) {
	var rs []*FileInfo
	if files, e := ioutil.ReadDir(path); e == nil {
		for _, file := range files {
			if file.IsDir() {
				rs = append(rs, list(tag, filepath.Join(path, file.Name()))...)
			} else {
				info := NewFileInfo(file, filepath.Join(path, file.Name()))
				info.AddTag(tag, path)
				rs = append(rs, info)
			}
		}
	}
	return rs
}
