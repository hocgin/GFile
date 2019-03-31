package file

import (
	"Web/src/model"
	provider "Web/src/provider/local"
)

func QueryAll() ([]*model.FileInfo) {
	all := provider.QueryAll()
	for _, v := range all {
		 v.IfLocalToRemote()
	}
	return all
}

func Page(page int, size int, classify string) ([]*model.FileInfo) {
	all := provider.Page(page, size, classify)
	for _, v := range all {
		v.IfLocalToRemote()
	}
	return all
}
