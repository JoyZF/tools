package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dateCmd)
}

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "时间戳转日期，如果输入时间戳，将会将其专成日期格式，并自动写入粘贴板",
	Long: `时间戳转日期，如果输入时间戳，将会将其专成日期格式，并自动写入粘贴板，
支持的时间格式: YYYY-MM-DD HH:MM:SS, YYYY/MM/DD HH/MM/SS, YYYY-MM-DD, YYYY/MM/DD`,
	Run: func(cmd *cobra.Command, args []string) {
		var date string
		if len(args) > 0 {
			timestamp := args[0]
			t, err := strconv.Atoi(timestamp)
			if err != nil {
				fmt.Println(err)
				return
			}
			switch len(timestamp) {
			case 10:
				date = time.Unix(int64(t), 0).Format(dateTime)
			case 13:
				date = time.Unix(int64(t/1e3), 0).Format(dateTime)
			default:
				fmt.Println("时间戳格式不正确")
				return
			}
		} else {
			date = time.Now().Format(dateTime)
		}
		_ = utils.WriteToClipboard(date)
	}}
