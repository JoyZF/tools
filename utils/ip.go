package utils

import (
	"io"
	"net/http"
)

func GetIp() (string, error) {
	resp, err := http.Get("https://ipv4.netarm.com/")
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	content, _ := io.ReadAll(resp.Body)
	return string(content), nil
}
