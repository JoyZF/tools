package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "push",
	Short: "封装一些常用的git命令,目前先实现git add . && git commit -m 'xxx' && git push",
	Long:  "封装一些常用的git命令,目前先实现git add . && git commit -m 'xxx' && git push",
	Run: func(cmd *cobra.Command, args []string) {
		var commitInfo string
		gstOutput, _ := exec.Command("git", "status", "-sb").Output()
		if len(args) > 0 {
			commitInfo = args[0]
		} else {
			commitInfo = fmt.Sprintf("feat: %s", string(gstOutput))
		}
		fmt.Println(commitInfo)
		if addOutput, err := exec.Command("git", "add", ".").Output(); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(addOutput))
		}

		if commitOutput, err := exec.Command("git", "commit", "-m", fmt.Sprintf("\"%s\"", commitInfo)).Output(); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(commitOutput))
		}
		if pushOutput, err := exec.Command("git", "push").Output(); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(pushOutput))
		}
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)
}
