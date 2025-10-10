package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [name] [path] [command]",
	Short: "Add a new project to the dev server list",
	Long: `Add a new project to your dev server list. You can specify the project name,
path to the project directory, and the command to run the dev server.

If no arguments are provided, an interactive mode will be used to collect the information.

Examples:
  dev add zensight-fe /path/to/zensight-fe "npm run dev"
  dev add api-server /home/user/api "go run main.go"
  dev add frontend ./frontend "yarn start"
  dev add  # Interactive mode`,
	Args: cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// Interactive mode
			runInteractiveAdd()
		} else if len(args) == 3 {
			// Non-interactive mode
			runNonInteractiveAdd(cmd, args)
		} else {
			fmt.Println("Error: Please provide all three arguments (name, path, command) or use interactive mode with no arguments")
			os.Exit(1)
		}
	},
}

func runInteractiveAdd() {
	var questions = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is your project named?",
				Help:    "Enter a unique name for your project",
			},
			Validate: survey.Required,
		},
		{
			Name: "path",
			Prompt: &survey.Input{
				Message: "What is the path to your project directory?",
				Help:    "Enter the absolute or relative path to your project directory",
			},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); ok {
					if strings.TrimSpace(str) == "" {
						return fmt.Errorf("path cannot be empty")
					}
					// Validate path
					absPath, err := filepath.Abs(str)
					if err != nil {
						return fmt.Errorf("invalid path: %v", err)
					}
					// Check if directory exists
					if _, err := os.Stat(absPath); os.IsNotExist(err) {
						return fmt.Errorf("directory '%s' does not exist", absPath)
					}
				}
				return nil
			},
		},
		{
			Name: "command",
			Prompt: &survey.Input{
				Message: "What command should be used to start the dev server?",
				Help:    "Enter the command to run your development server (e.g., 'npm run dev', 'go run main.go', 'yarn start')",
			},
			Validate: survey.Required,
		},
		{
			Name: "description",
			Prompt: &survey.Input{
				Message: "What is the description for this project?",
				Help:    "Enter an optional description for your project (press Enter to skip)",
			},
		},
	}

	answers := struct {
		Name        string `survey:"name"`
		Path        string `survey:"path"`
		Command     string `survey:"command"`
		Description string `survey:"description"`
	}{}

	// Perform the questions
	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Printf("Error during interactive setup: %v\n", err)
		os.Exit(1)
	}

	// Convert path to absolute path
	absPath, err := filepath.Abs(answers.Path)
	if err != nil {
		fmt.Printf("Error: Invalid path '%s': %v\n", answers.Path, err)
		os.Exit(1)
	}

	// Add the project
	if err := storage.AddProject(answers.Name, absPath, answers.Command, answers.Description); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Successfully added project '%s'\n", answers.Name)
	fmt.Printf("   Path: %s\n", absPath)
	fmt.Printf("   Command: %s\n", answers.Command)
	if answers.Description != "" {
		fmt.Printf("   Description: %s\n", answers.Description)
	}
}

func runNonInteractiveAdd(cmd *cobra.Command, args []string) {
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
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("description", "d", "", "Description for the project")
}
