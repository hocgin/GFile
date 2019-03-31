package result

import "github.com/gin-gonic/gin/json"

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) (*Result) {
	return &Result{Code: 200, Message: "ok", Data: data}
}

func Error(code int, message string) (*Result) {
	return &Result{Code: code, Message: message, Data: nil}
}

// 转字节数组
func (r *Result) ToByte() ([]byte) {
	bytes, _ := json.Marshal(r)
	return bytes
}
