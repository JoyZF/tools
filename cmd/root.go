package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "tools is a command line tool",
	Long:  "tools is a command line tool, detail for https://github.com/JoyZF/tools/README.md",
}

func init() {

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
