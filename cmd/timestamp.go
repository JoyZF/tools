package cmd

import (
	"errors"
	"fmt"
	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(timestampCmd)
}

// timestampCmd represents date to timestamp
var timestampCmd = &cobra.Command{
	Use:   "time",
	Short: "获取当前时间戳 如果传递时间则转化对应的时间为时间戳, 并自动写入粘贴板",
	Long: `获取当前时间戳 如果传递时间则转化对应的时间为时间戳, 并自动写入粘贴板
           支持的时间格式: YYYY-MM-DD HH:MM:SS, YYYY/MM/DD HH/MM/SS, YYYY-MM-DD, YYYY/MM/DD`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			timestamp int64
			err       error
		)

		if len(args) > 0 {
			if timestamp, err = value2Timestamp(args[0]); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			timestamp = time.Now().Unix()
		}

		_ = utils.WriteToClipboard(fmt.Sprintf("%d", timestamp))
	},
}

func value2Timestamp(val string) (int64, error) {
	loc, err := time.ParseInLocation(dateTime, val, time.Local)
	if err == nil {
		return loc.Unix(), nil
	}
	loc, err = time.ParseInLocation(dateTimeF, val, time.Local)
	if err == nil {
		return loc.Unix(), nil
	}

	loc, err = time.ParseInLocation(date, val, time.Local)
	if err == nil {
		return loc.Unix(), nil
	}
	loc, err = time.ParseInLocation(dateF, val, time.Local)
	if err == nil {
		return loc.Unix(), nil
	}
	return 0, errors.New("val is not time format")
}
