package utils

import (
	"io"
	"net/http"
)

func GetKFC() (string, error) {
	resp, err := http.Get("http://shanhe.kim/api/za/kfc.php")
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	content, _ := io.ReadAll(resp.Body)
	return string(content), nil
}
