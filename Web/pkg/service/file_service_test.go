package service

import (
	"GFile/pkg/util"
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	escape := url.QueryEscape("sd/wes/cjhsd")
	t.Log(escape)
}

func TestIP(t *testing.T) {
	fmt.Println(util.GetOutboundIP())
}

