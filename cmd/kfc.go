package cmd

import (
	"fmt"
	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
	"os/exec"
)

var kfcCmd = &cobra.Command{
	Use:   "kfc",
	Short: "随机返回疯狂星期四文案",
	Long:  "随机返回疯狂星期四文案",
	Run: func(cmd *cobra.Command, args []string) {
		kfc, err := utils.GetKFC()
		if err != nil {
			fmt.Println(err)
			return
		}

		output, _ := exec.Command("echo", kfc).Output()
		fmt.Println(string(output))
		_ = utils.WriteToClipboard(kfc)
	},
}

func init() {
	rootCmd.AddCommand(kfcCmd)
}
