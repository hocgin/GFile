package model

import (
	"Web/src/config"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

// typeOf
const (
	VIDEO   = "video"
	AUDIO   = "audio"
	TEXT    = "text"
	IMAGE   = "image"
	UNKNOWN = "unknown"
)

// source
const (
	LOCAL = "local"
	URL   = "url"
)

type FileInfo struct {
	//, 分隔
	Tags []string `json:"tags"`
	// 类型
	TypeOf string `json:"typeOf"`
	// 存储类型
	Source string `json:"source"`
	// 文件大小
	Size int64 `json:"size"`
	// 文件名
	FileName string `json:"fileName"`
	// path
	Path string `json:"path"`
	// 变更时间
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewFileInfo(info os.FileInfo, path string) (*FileInfo) {
	rs := new(FileInfo)
	rs.Size = info.Size()
	rs.FileName = info.Name()
	rs.Source = LOCAL
	rs.TypeOf = getFileType(rs.FileName)
	rs.Path = path
	rs.UpdatedAt = info.ModTime()
	rs.Tags = []string{getFileSuffix(rs.FileName)}
	return rs
}

// 获取文件后缀
func getFileSuffix(fileName string) string {
	if strings.Contains(fileName, ".") {
		arr := strings.Split(fileName, ".")
		return arr[len(arr)-1]
	}
	return ""
}

// get file type from analysis file name suffix
func getFileType(fileName string) string {
	switch getFileSuffix(fileName) {
	case "mp3", "mp4":
		return VIDEO
	case "aif":
		return AUDIO
	case "txt", "md":
		return TEXT
	case "png":
		return IMAGE
	}
	return UNKNOWN
}

func (that *FileInfo) IfLocalToRemote() {
	switch that.Source {
	case LOCAL:
		switch that.TypeOf {
		case VIDEO: // 视频
			path := that.Path
			u := strings.ReplaceAll(constant.VIDEO_ANALYSIS_URL, "{:id}", url.QueryEscape(path))
			// eg: http://127.0.0.1:8080/analysis/video?id=xxx
			that.Source = URL
			that.Path = fmt.Sprintf("%s%s", constant.GetHostName(), u)
		default:
			path := that.Path
			u := strings.ReplaceAll(constant.FILE_ANALYSIS_URL, "{:id}", url.QueryEscape(path))
			that.Path = fmt.Sprintf("%s%s", constant.GetHostName(), u)
		}
	default:
	}
}

// 数据提供者
// 1. 来自爬虫整理数据
// 2. 来自硬盘读取数据
