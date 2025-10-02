package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a project from the dev server list",
	Long: `Remove a project from your registered projects. This will permanently
delete the project configuration.

Examples:
  dev remove zensight-fe
  dev remove api-server`,
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
		
		// Confirm removal
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			fmt.Printf("Are you sure you want to remove project '%s'? (y/N): ", name)
			var response string
			fmt.Scanln(&response)
			if response != "y" && response != "Y" && response != "yes" {
				fmt.Println("Operation cancelled.")
				return
			}
		}
		
		// Remove the project
		if err := storage.RemoveProject(name); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("✅ Successfully removed project '%s'\n", name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolP("force", "f", false, "Remove without confirmation")
}
