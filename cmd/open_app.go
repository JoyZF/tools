package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var supportAppMap = map[string]string{
	"c": "Google Chrome.app",
	"w": "WeChat.app",
	"n": "Navicat Premium.app",
	"s": "Sublime Text.app",
	"t": "Typora.app",
}

var openAppCmd = &cobra.Command{
	Use:   "open",
	Short: "打开应用程序",
	Long:  "打开应用程序, detail for https://github.com/JoyZF/tools/README.md",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			_ = cmd.Help()
			return
		}
		if _, ok := supportAppMap[name]; !ok {
			fmt.Println(supportList())
			return
		}
		fmt.Println(openApp(supportAppMap[name]))
	},
}

func init() {
	rootCmd.AddCommand(openAppCmd)
	openAppCmd.Flags().StringP("name",
		"n", "Google\\ Chrome.app",
		`支持指定程序快速打开,
				目前支持：
					c: Google\ Chrome.app
					w: WeChat.app
					n: Navicat\ Premium.app
					s: Sublime\ Text.app
					t: Typora.app
`)
}

func openApp(app string) error {
	return exec.Command("open", fmt.Sprintf("/Applications/%s", app)).Run()
}

func supportList() string {
	var str string
	for k, v := range supportAppMap {
		str = fmt.Sprintf("关键词: %s 应用程序: %s\n", k, v)
	}
	return str
}
