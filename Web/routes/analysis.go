package routes

import (
	"GFile/pkg/analysis"
	"GFile/pkg/core"
	"net/http"
)

// analysis get video stream
func AnalysisVideo(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok && len(id[0]) > 1 {
		if path, ok := analysis.GetFilePath(id[0]); ok {
			http.ServeFile(w, r, path)
		}
	}

	write(w, r, core.Error(500, "error"))
}

func AnalysisFile(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok && len(id[0]) > 1 {
		if path, ok := analysis.GetFilePath(id[0]); ok {
			http.ServeFile(w, r, path)
		}
	}
	write(w, r, core.Error(500, "error"))
}
