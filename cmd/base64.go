package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/JoyZF/tools/utils"
	"github.com/spf13/cobra"
	"os/exec"
)

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64 encode or decode",
	Long:  `base64 encode or decode, 第二个参数为encode或decode，第三个参数为需要转换的字符串`,
	Run: func(cmd *cobra.Command, args []string) {
		do, err := cmd.Flags().GetString("do")
		val, err := cmd.Flags().GetString("string")
		if err != nil {
			_ = cmd.Help()
			return
		}
		var res string
		switch {
		case do == "encode" || do == "e":
			res = encode(val)
		case do == "decode" || do == "d":
			res = decode(val)
		}
		output, _ := exec.Command("echo", fmt.Sprintf("raw is %s res is %s", val, res)).Output()
		fmt.Println(string(output))
		_ = utils.WriteToClipboard(res)
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringP("do", "d", "encode", "可选参数，decode或encode 或者d或e")
	base64Cmd.Flags().StringP("string", "s", "", "需要转换的字符串")
}

func encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func decode(value string) string {
	bytes, _ := base64.StdEncoding.DecodeString(value)
	return string(bytes)
}
