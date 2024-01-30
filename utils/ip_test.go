package utils

import (
	"fmt"
	"testing"
)

func TestGetIp(t *testing.T) {
	ip, _ := GetIp()
	fmt.Println(ip)
}
