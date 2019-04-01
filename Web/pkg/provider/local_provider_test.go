package provider

import (
	"strings"
	"testing"
)

func TestQueryAll(t *testing.T) {
	//localProvider := LocalProvider{}
	//infos, _ := localProvider.li()
	//t.Log(infos)
}

func TestA(t *testing.T) {
	all := strings.Split("sd.sd.mp3", ".")

	t.Log(all[len(all)-1])
}