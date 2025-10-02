package cmd

import (
	"dev-util/storage"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered projects",
	Long: `List all registered projects with their details including name, path,
command, and creation date.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := storage.ListProjects()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
		if len(projects) == 0 {
			fmt.Println("No projects registered. Use 'dev add' to add your first project.")
			return
		}
		
		fmt.Printf("📋 Registered Projects (%d total)\n\n", len(projects))
		
		// Create tabwriter for aligned output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tPATH\tCOMMAND\tCREATED")
		fmt.Fprintln(w, "----\t----\t-------\t-------")
		
		for _, project := range projects {
			created := project.CreatedAt.Format("2006-01-02")
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", 
				project.Name, 
				project.Path, 
				project.Command,
				created)
		}
		
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
