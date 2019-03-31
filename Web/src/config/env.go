package constant

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	HOSTNAME = "http://{:ip}:{:port}"
)

func GetHostName() string {
	port := viper.GetString(SERVER_PORT)
	ip := viper.GetString(SERVER_IP)
	return strings.ReplaceAll(strings.ReplaceAll(HOSTNAME, "{:ip}", ip), "{:port}", port)
}