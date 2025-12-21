package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd [name]",
	Short: "Get the path to a project directory",
	Long: `Get the path to a registered project's directory.

This command is designed to work with shell integration. To set it up:
  eval "$(dev init bash)"   # for bash
  eval "$(dev init zsh)"    # for zsh
  dev init fish | source    # for fish

Then use 'dev-cd <project>' to change directories.

Examples:
  dev-cd zensight-fe
  dev-cd api-server`,
	Args: cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		projects, err := storage.ListProjects()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var projectNames []string
		for _, project := range projects {
			if toComplete == "" || len(project.Name) >= len(toComplete) && project.Name[:len(toComplete)] == toComplete {
				projectNames = append(projectNames, project.Name)
			}
		}

		return projectNames, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		project, err := storage.GetProject(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if _, err := os.Stat(project.Path); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Project directory '%s' no longer exists\n", project.Path)
			os.Exit(1)
		}

		pathOnly, _ := cmd.Flags().GetBool("path")
		if pathOnly {
			fmt.Print(project.Path)
			return
		}

		fmt.Printf("📁 Project '%s' information:\n", name)
		fmt.Printf("   Path: %s\n", project.Path)
		if project.Description != "" {
			fmt.Printf("   Description: %s\n", project.Description)
		}
		fmt.Println()
		fmt.Println("To change to this directory, set up shell integration:")
		fmt.Println()
		fmt.Println("  eval \"$(dev init bash)\"   # for bash")
		fmt.Println("  eval \"$(dev init zsh)\"    # for zsh")
		fmt.Println("  dev init fish | source    # for fish")
		fmt.Println()
		fmt.Println("Then use: dev-cd", name)
	},
}

func init() {
	cdCmd.Flags().Bool("path", false, "Output only the project path (for shell integration)")
	rootCmd.AddCommand(cdCmd)
}
