package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new project",
	Long: `Create a new project with the specified name.

The name parameter is required.`,
	Args: cobra.ExactArgs(1),
	Run:  RunNew,
}

func RunNew(cmd *cobra.Command, args []string) {
	name := args[0]

	// Get rocket root
	rocketRoot := GetRocketRoot()

	// Create the project directory under rocket root
	projectPath := filepath.Join(rocketRoot, name)
	err := os.MkdirAll(projectPath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", projectPath, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Created new project directory: %s\n", projectPath)
}

func init() {
	rootCmd.AddCommand(newCmd)
}
