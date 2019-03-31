package local

import (
	"Web/src/config"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"testing"
)

func TestQueryAll(t *testing.T) {
	//localProvider := LocalProvider{}
	//infos, _ := localProvider.li()
	//t.Log(infos)
}

func TestQ(t *testing.T) {
	infos := list("/Users/hocgin/Downloads")
	t.Log(infos)
}

func TestA(t *testing.T) {
	all := strings.Split("sd.sd.mp3", ".")

	t.Log(all[len(all)-1])
}
func TestConfig(t *testing.T) {
	config := viper.New()
	config.SetConfigName("application")
	config.AddConfigPath("/Users/hocgin/Document/Projects/Go/src/Web/config/")
	config.SetConfigType("yaml")
	if e := config.ReadInConfig(); e != nil {
		fmt.Println(e.Error())
	}
	i := config.GetString(constant.FILE_PATH)
	fmt.Println(i)
}

func TestAnalysisLocalFile(t *testing.T) {

}