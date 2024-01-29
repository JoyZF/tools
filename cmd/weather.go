package cmd

import (
	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "查询天气, 默认查询当前城市, 支持输入城市查询",
	Long:  "查询天气, 默认查询当前城市, 支持输入城市查询",
	Run: func(cmd *cobra.Command, args []string) {
		weather := utils.Weather{}
		if len(args) == 0 {
			weather.GetWeather("")
		} else {
			weather.GetWeather(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
