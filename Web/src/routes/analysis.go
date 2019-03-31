package routes

import (
	"Web/src/core/result"
	"Web/src/service/analysis"
	"net/http"
)

// analysis get video stream
func AnalysisVideo(w http.ResponseWriter, r *http.Request)  {
	if id, ok := r.URL.Query()["id"]; ok && len(id[0]) > 1 {
		if path, ok := analysis.GetFilePath(id[0]); ok {
			http.ServeFile(w, r, path)
		}
	}

	write(w, r, result.Error(500, "error"))
}

func AnalysisFile(w http.ResponseWriter, r *http.Request)  {
	if id, ok := r.URL.Query()["id"]; ok && len(id[0]) > 1 {
		if path, ok := analysis.GetFilePath(id[0]); ok {
			http.ServeFile(w, r, path)
		}
	}
	write(w, r, result.Error(500, "error"))
}