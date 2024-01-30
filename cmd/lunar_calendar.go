package cmd

import (
	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
)

var lunarCalendarCmd = &cobra.Command{
	Use:   "luck",
	Short: "返回今日老黄历",
	Long:  "返回今日老黄历",
	Run: func(cmd *cobra.Command, args []string) {
		utils.LunarCalendarTable()
	},
}

func init() {
	rootCmd.AddCommand(lunarCalendarCmd)
}
