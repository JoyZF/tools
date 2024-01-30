package cmd

import (
	"fmt"
	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
	"os/exec"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "获取当前IP,并写入粘贴板",
	Long:  "获取当前IP,并写入粘贴板",
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := utils.GetIp()
		if err != nil {
			fmt.Println(err)
			return
		}

		output, _ := exec.Command("echo", ip).Output()
		fmt.Println(string(output))
		_ = utils.WriteToClipboard(ip)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
