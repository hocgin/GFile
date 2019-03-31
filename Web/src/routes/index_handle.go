package routes

import (
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["time"] = time.Now()

	render(w, "index", data)
}
