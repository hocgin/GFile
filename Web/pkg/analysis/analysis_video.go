package analysis

import (
	"GFile/pkg/provider"
	"GFile/pkg/util"
)

// 获取文件路径
func GetFilePath(id string) (string, bool) {
	if path, ok := provider.Decode(id); ok {
		if util.IsExists(path) {
			return path, true
		}
	}
	return "", false
}

