package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name] [path] [command]",
	Short: "Add a new project to the dev server list",
	Long: `Add a new project to your dev server list. You can specify the project name,
path to the project directory, and the command to run the dev server.

Examples:
  dev add zensight-fe /path/to/zensight-fe "npm run dev"
  dev add api-server /home/user/api "go run main.go"
  dev add frontend ./frontend "yarn start"`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		path := args[1]
		command := args[2]
		
		// Validate path
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error: Invalid path '%s': %v\n", path, err)
			os.Exit(1)
		}
		
		// Check if directory exists
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			fmt.Printf("Error: Directory '%s' does not exist\n", absPath)
			os.Exit(1)
		}
		
		// Get description from flag if provided
		description, _ := cmd.Flags().GetString("description")
		
		// Add the project
		if err := storage.AddProject(name, absPath, command, description); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("✅ Successfully added project '%s'\n", name)
		fmt.Printf("   Path: %s\n", absPath)
		fmt.Printf("   Command: %s\n", command)
		if description != "" {
			fmt.Printf("   Description: %s\n", description)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("description", "d", "", "Description for the project")
}
