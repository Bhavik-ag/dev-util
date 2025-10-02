package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev",
	Short: "A CLI tool to manage and start dev servers for your projects",
	Long: `Dev is a CLI tool that helps you manage and start development servers 
for your projects from anywhere. You can register projects and start their 
dev servers with simple commands.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
