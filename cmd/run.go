package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Start the dev server for a project",
	Long: `Start the development server for a registered project. The command will
change to the project directory and execute the configured command.

Examples:
  dev run zensight-fe
  dev run api-server`,
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
		
		fmt.Printf("🚀 Starting dev server for '%s'...\n", name)
		fmt.Printf("   Path: %s\n", project.Path)
		fmt.Printf("   Command: %s\n", project.Command)
		if project.Description != "" {
			fmt.Printf("   Description: %s\n", project.Description)
		}
		fmt.Println()
		
		// Parse command and arguments
		parts := strings.Fields(project.Command)
		if len(parts) == 0 {
			fmt.Printf("Error: Invalid command '%s'\n", project.Command)
			os.Exit(1)
		}
		
		// Create command
		execCmd := exec.Command(parts[0], parts[1:]...)
		execCmd.Dir = project.Path
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin
		
		// Start the command
		if err := execCmd.Run(); err != nil {
			fmt.Printf("Error running command: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
