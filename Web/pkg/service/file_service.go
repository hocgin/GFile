package service

import (
	"GFile/pkg/provider"
)

func QueryAll() ([]*provider.FileInfo) {
	all := provider.QueryAll()
	for _, v := range all {
		v.IfLocalToRemote()
	}
	return all
}

func Page(page int, size int, classify string) ([]*provider.FileInfo) {
	all := provider.Page(page, size, classify)
	for _, v := range all {
		v.IfLocalToRemote()
	}
	return all
}
