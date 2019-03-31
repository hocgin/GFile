package analysis

import (
	"Web/src/provider/local"
	"Web/src/util"
)

// 获取文件路径
func GetFilePath(id string) (string, bool) {
	if path, ok := local.Decode(id); ok {
		if util.IsExists(path) {
			return path, true
		}
	}
	return "", false
}

