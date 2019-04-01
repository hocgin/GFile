package constant

import (
	"GFile/pkg/config"
	"github.com/spf13/viper"
	"strings"
)

const (
	HOSTNAME = "http://{:ip}:{:port}"
)

func GetHostName() string {
	port := viper.GetString(config.SERVER_PORT)
	ip := viper.GetString(config.SERVER_IP)
	return strings.ReplaceAll(strings.ReplaceAll(HOSTNAME, "{:ip}", ip), "{:port}", port)
}