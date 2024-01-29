package utils

import (
	"fmt"
	"testing"
)

func TestWeather_GetWeather(t *testing.T) {
	fmt.Println(getWeatherInfo("西宁"))
}

func TestWeather_GetWeather1(t *testing.T) {
	w := Weather{}
	w.GetWeather("杭州")
}
