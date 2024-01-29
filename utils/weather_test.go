package utils

import (
	"fmt"
	"testing"
)

func TestWeather_GetWeather(t *testing.T) {
	fmt.Println(getWeatherInfo("西宁"))
}
