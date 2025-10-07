package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd [name]",
	Short: "Change directory to a project",
	Long: `Change directory to a registered project. This command will navigate
to the project's directory and start a new shell session in that location.

Examples:
  dev cd zensight-fe
  dev cd api-server`,
	Args: cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Get list of registered projects for completion
		projects, err := storage.ListProjects()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		
		var projectNames []string
		for _, project := range projects {
			// Only show projects that match the current input
			if toComplete == "" || len(project.Name) >= len(toComplete) && project.Name[:len(toComplete)] == toComplete {
				projectNames = append(projectNames, project.Name)
			}
		}
		
		// Return with NoSpace directive to prevent adding space after completion
		return projectNames, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		
		// Get project details
		project, err := storage.GetProject(name)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
		// Check if directory still exists
		if _, err := os.Stat(project.Path); os.IsNotExist(err) {
			fmt.Printf("Error: Project directory '%s' no longer exists\n", project.Path)
			os.Exit(1)
		}
		
		fmt.Printf("📁 Changing to project '%s'...\n", name)
		fmt.Printf("   Path: %s\n", project.Path)
		if project.Description != "" {
			fmt.Printf("   Description: %s\n", project.Description)
		}
		fmt.Println()
		
		// Start a new shell in the project directory
		var shellCmd *exec.Cmd
		
		if runtime.GOOS == "windows" {
			// On Windows, use cmd.exe
			shellCmd = exec.Command("cmd")
		} else {
			// On Unix-like systems, use the current shell or default to bash
			shell := os.Getenv("SHELL")
			if shell == "" {
				shell = "/bin/bash"
			}
			shellCmd = exec.Command(shell)
		}
		
		// Set the working directory to the project path
		shellCmd.Dir = project.Path
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		shellCmd.Stdin = os.Stdin
		
		// Start the shell
		if err := shellCmd.Run(); err != nil {
			fmt.Printf("Error starting shell: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
